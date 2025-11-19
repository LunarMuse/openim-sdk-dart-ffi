//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	soName  = "libopenimsdk"                     // 库名称
	rootOutPath = "../shared/"                   // 根输出路径，避免+=覆盖
	goSrc   = "go"                               // Go源码目录
	minIOSVersion = "13.0"                       // 最低支持iOS版本
)

var Default = Build

// BuildAll compiles the project for all platforms.
func Build() {
	platforms := []struct {
		name string
		fn   func() error
	}{
		{"Android", BuildAndroid},
		{"iOS", BuildIOS},
		{"Linux", BuildLinux},
		{"Windows", BuildWindows},
		{"MacOS", BuildMacOS},
	}

	for _, p := range platforms {
		fmt.Printf("=== Building for %s ===\n", p.name)
		if err := p.fn(); err != nil {
			fmt.Printf("Error building for %s: %v\n", p.name, err)
		} else {
			fmt.Printf("Successfully built for %s\n", p.name)
		}
	}
}

// -------------------------- Android 编译逻辑（保留原有） --------------------------
func buildAndroid(aOutPath, arch, apiLevel string) error {
	fmt.Printf("Building for Android %s...\n", arch)

	ndkPath := os.Getenv("ANDROID_NDK_HOME")
	if ndkPath == "" {
		return fmt.Errorf("ANDROID_NDK_HOME environment variable not set")
	}

	osSuffix := ""
	if runtime.GOOS == "windows" {
		osSuffix = ".cmd"
	}

	ccBasePath := filepath.Join(ndkPath, "toolchains", "llvm", "prebuilt", runtime.GOOS+"-x86_64", "bin")
	if _, err := os.Stat(ccBasePath); err != nil {
		return fmt.Errorf("NDK toolchain path not found: %s", ccBasePath)
	}

	var cc string
	switch arch {
	case "arm":
		cc = filepath.Join(ccBasePath, "armv7a-linux-androideabi"+apiLevel+"-clang"+osSuffix)
	case "arm64":
		cc = filepath.Join(ccBasePath, "aarch64-linux-android"+apiLevel+"-clang"+osSuffix)
	case "386":
		cc = filepath.Join(ccBasePath, "i686-linux-android"+apiLevel+"-clang"+osSuffix)
	case "amd64":
		cc = filepath.Join(ccBasePath, "x86_64-linux-android"+apiLevel+"-clang"+osSuffix)
	default:
		return fmt.Errorf("unsupported Android arch: %s", arch)
	}

	// 构建输出目录
	if err := os.MkdirAll(filepath.Join(aOutPath, arch), 0755); err != nil {
		return err
	}

	env := []string{
		"CGO_ENABLED=1",
		"GOOS=android",
		"GOARCH=" + arch,
		"CC=" + cc,
	}
	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", filepath.Join(aOutPath, arch, soName+".so"), ".")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Dir = goSrc
	cmd.Env = append(os.Environ(), env...)
	return cmd.Run()
}

func BuildAndroid() error {
	androidOutPath := filepath.Join(rootOutPath, "android")
	architectures := []struct {
		Arch, API string
	}{
		{"arm", "16"},
		{"arm64", "21"},
		{"386", "16"},
		{"amd64", "21"},
	}

	for _, arch := range architectures {
		if err := buildAndroid(androidOutPath, arch.Arch, arch.API); err != nil {
			fmt.Printf("Failed to build for Android %s: %v\n", arch.Arch, err)
		}
	}
	return nil
}

// -------------------------- iOS 编译逻辑（核心修复） --------------------------
// 获取iOS SDK路径
func getIOSSDKPath(sdk string) (string, error) {
	cmd := exec.Command("xcrun", "--sdk", sdk, "--show-sdk-path")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get %s SDK path: %w\n%s", sdk, err, string(output))
	}
	return strings.TrimSpace(string(output)), nil
}

// 获取iOS Clang编译器路径
func getIOSCC(sdk string) (string, error) {
	cmd := exec.Command("xcrun", "--sdk", sdk, "-f", "clang")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get %s clang path: %w\n%s", sdk, err, string(output))
	}
	return strings.TrimSpace(string(output)), nil // 关键：Trim空格避免路径错误
}

// 编译单个iOS架构的静态库
func buildIOSArch(arch, sdk, minOS, output string) error {
	// 获取SDK路径和编译器
	sdkPath, err := getIOSSDKPath(sdk)
	if err != nil {
		return err
	}
	cc, err := getIOSCC(sdk)
	if err != nil {
		return err
	}

	// 构建CGO编译标志
	var cgoArch, sdkName string
	switch sdk {
	case "iphoneos":
		sdkName = "iphoneos"
		cgoArch = arch // 真机arm64
	case "iphonesimulator":
		sdkName = "iphonesimulator"
		cgoArch = arch // 模拟器x86_64/arm64
	default:
		return fmt.Errorf("unsupported iOS SDK: %s", sdk)
	}

	// CGO编译参数：指定架构、SDK根目录、最低iOS版本
	cflags := fmt.Sprintf("-arch %s -isysroot %s -m%s-version-min=%s",
		cgoArch, sdkPath, sdkName, minOS)

	// 构建环境变量（局部传递，不全局覆盖）
	env := []string{
		"CGO_ENABLED=1",
		"GOOS=ios", // 关键：iOS的GOOS设置为ios（Go 1.16+支持）
		"GOARCH=" + arch,
		"CC=" + cc,
		"CGO_CFLAGS=" + cflags,
		"CGO_LDFLAGS=" + cflags,
	}

	fmt.Printf("Building iOS %s (%s) library...\n", arch, sdk)
	cmd := exec.Command("go", "build", "-buildmode=c-archive", "-trimpath", "-ldflags=-s -w", "-o", output, ".")
	cmd.Dir = goSrc
	cmd.Env = append(os.Environ(), env...) // 合并系统环境变量
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build iOS %s: %w", arch, err)
	}
	return nil
}

// 合并多个架构为通用胖库
func createIOSUniversalLibrary(inputs []string, output string) error {
	fmt.Println("Creating iOS universal library...")
	if _, err := exec.LookPath("lipo"); err != nil {
		return fmt.Errorf("lipo tool not found (install Xcode): %w", err)
	}

	// 构建lipo命令：-create 输入1 输入2 -output 输出
	args := append([]string{"-create"}, inputs...)
	args = append(args, "-output", output)
	cmd := exec.Command("lipo", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("lipo failed: %w", err)
	}

	// 可选：删除临时架构库
	for _, input := range inputs {
		_ = os.Remove(input)
		_ = os.Remove(input[:len(input)-2] + ".h") // 删除对应的头文件
	}
	return nil
}

func BuildIOS() error {
	iosOutPath := filepath.Join(rootOutPath, "ios")
	if err := os.MkdirAll(iosOutPath, 0755); err != nil {
		return err
	}

	// 定义需要编译的iOS架构
	architectures := []struct {
		Arch string // GOARCH
		SDK  string // iphoneos/iphonesimulator
	}{
		{"arm64", "iphoneos"},          // 真机arm64
		{"arm64", "iphonesimulator"},   // Apple Silicon Mac模拟器（可选，按需开启）
	}

	// 存储临时架构库路径
	tempLibs := make([]string, 0, len(architectures))
	for _, arch := range architectures {
		// 临时输出路径：如 libopenimsdk_arm64_iphoneos.a
		tempOutput := filepath.Join(iosOutPath, fmt.Sprintf("%s_%s_%s.a", soName, arch.Arch, arch.SDK))
		if err := buildIOSArch(arch.Arch, arch.SDK, minIOSVersion, tempOutput); err != nil {
			fmt.Printf("Failed to build iOS %s (%s): %v\n", arch.Arch, arch.SDK, err)
			// 可选择继续编译其他架构或直接返回错误
			// return err
		} else {
			tempLibs = append(tempLibs, tempOutput)
		}
	}

	// 合并为通用库（仅当有多个架构时）
	if len(tempLibs) > 0 {
		finalOutput := filepath.Join(iosOutPath, soName+".a")
		if err := createIOSUniversalLibrary(tempLibs, finalOutput); err != nil {
			return err
		}
		fmt.Printf("iOS universal library built at: %s\n", finalOutput)
	} else {
		return fmt.Errorf("no iOS architectures were built successfully")
	}

	return nil
}

// -------------------------- MacOS 编译逻辑（修复路径） --------------------------
func BuildMacOS() error {
	fmt.Println("Building for MacOS...")
	macosOutPath := filepath.Join(rootOutPath, "macos")
	if err := os.MkdirAll(macosOutPath, 0755); err != nil {
		return err
	}

	arch := os.Getenv("GOARCH")
	if arch == "" {
		arch = runtime.GOARCH
	}

	// 局部环境变量，不全局覆盖
	env := []string{
		"CGO_ENABLED=1",
		"GOOS=darwin",
		"GOARCH=" + arch,
		"CC=clang",
	}

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", filepath.Join(macosOutPath, soName+".dylib"), ".")
	cmd.Dir = goSrc
	cmd.Env = append(os.Environ(), env...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build MacOS: %w", err)
	}
	return nil
}

// -------------------------- Linux 编译逻辑（修复路径） --------------------------
func BuildLinux() error {
	fmt.Println("Building for Linux...")
	linuxOutPath := filepath.Join(rootOutPath, "linux")
	if err := os.MkdirAll(linuxOutPath, 0755); err != nil {
		return err
	}

	arch := os.Getenv("GOARCH")
	if arch == "" {
		arch = runtime.GOARCH
	}
	cc := os.Getenv("CC")
	if cc == "" {
		cc = "gcc"
	}

	env := []string{
		"CGO_ENABLED=1",
		"GOOS=linux",
		"GOARCH=" + arch,
		"CC=" + cc,
	}
	if cxx := os.Getenv("CXX"); cxx != "" {
		env = append(env, "CXX="+cxx)
	}

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", filepath.Join(linuxOutPath, soName+".so"), ".")
	cmd.Dir = goSrc
	cmd.Env = append(os.Environ(), env...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build Linux: %w", err)
	}
	return nil
}

// -------------------------- Windows 编译逻辑（修复路径） --------------------------
func BuildWindows() error {
	fmt.Println("Building for Windows...")
	windowsOutPath := filepath.Join(rootOutPath, "windows")
	if err := os.MkdirAll(windowsOutPath, 0755); err != nil {
		return err
	}

	arch := os.Getenv("GOARCH")
	if arch == "" {
		arch = runtime.GOARCH
	}
	cc := os.Getenv("CC")
	if cc == "" {
		cc = "gcc" // 需安装MinGW-w64
	}

	env := []string{
		"CGO_ENABLED=1",
		"GOOS=windows",
		"GOARCH=" + arch,
		"CC=" + cc,
	}
	if cxx := os.Getenv("CXX"); cxx != "" {
		env = append(env, "CXX="+cxx)
	}

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", filepath.Join(windowsOutPath, soName+".dll"), ".")
	cmd.Dir = goSrc
	cmd.Env = append(os.Environ(), env...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build Windows: %w", err)
	}
	return nil
}
