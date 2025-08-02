//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	soName  = "libopenimsdk" //
	outPath = "../shared/"
	goSrc   = "go" //
)

var Default = Build

// BuildAll compiles the project for all platforms.
func Build() {
	if err := BuildAndroid(); err != nil {
		fmt.Println("Error building for Android:", err)
	}
	if err := BuildIOS(); err != nil {
		fmt.Println("Error building for iOS:", err)
	}
	if err := BuildLinux(); err != nil {
		fmt.Println("Error building for Linux:", err)
	}
	if err := BuildWindows(); err != nil {
		fmt.Println("Error building for Windows:", err)
	}
	if err := BuildMacOS(); err != nil {
		fmt.Println("Error building for MacOS:", err)
	}
}

func buildAndroid(aOutPath, arch, apiLevel string) error {
	fmt.Printf("Building for %s...\n", arch)

	ndkPath := os.Getenv("ANDROID_NDK_HOME")
	osSuffix := ""
	if runtime.GOOS == "windows" {
		osSuffix = ".cmd" //
	}

	ccBasePath := ndkPath + "/toolchains/llvm/prebuilt/" + runtime.GOOS + "-x86_64/bin/"

	var cc string
	switch arch {
	case "arm":
		cc = ccBasePath + "armv7a-linux-androideabi" + apiLevel + "-clang" + osSuffix
	case "arm64":
		cc = ccBasePath + "aarch64-linux-android" + apiLevel + "-clang" + osSuffix
	case "386":
		cc = ccBasePath + "i686-linux-android" + apiLevel + "-clang" + osSuffix
	case "amd64":
		cc = ccBasePath + "x86_64-linux-android" + apiLevel + "-clang" + osSuffix
	}

	env := []string{
		"CGO_ENABLED=1",
		"GOOS=android",
		"GOARCH=" + arch,
		"CC=" + cc,
	}
	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", aOutPath+"/"+arch+"/"+soName+".so", ".")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Dir = goSrc
	cmd.Env = append(os.Environ(), env...)
	return cmd.Run()
}

// BuildAndroid compiles the project for Android.
func BuildAndroid() error {
	architectures := []struct {
		Arch, API string
	}{
		{"arm", "16"},
		{"arm64", "21"},
		{"386", "16"},
		{"amd64", "21"},
	}

	for _, arch := range architectures {
		if err := buildAndroid(outPath+"android", arch.Arch, arch.API); err != nil {
			fmt.Printf("Failed to build for %s: %v\n", arch.Arch, err)
		}
	}
	return nil
}

// BuildMacOS compiles the project for MacOS.
func BuildMacOS() error {
	fmt.Println("Building for MacOS...")
	outPath += "macos"
	arch := os.Getenv("GOARCH")

	if len(arch) == 0 {
		arch = runtime.GOARCH
	}

	os.Setenv("GOOS", "darwin")
	os.Setenv("GOARCH", arch)
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", "clang")

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-o", outPath+"/"+soName+".dylib", ".")
	cmd.Dir = goSrc
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to build for MacOS: %v\n", err)
		return err
	}
	fmt.Println("Build for MacOS completed successfully.")
	return nil
}

func getIOSSDKPath(sdk string) (string, error) {
	cmd := exec.Command("xcrun", "--sdk", sdk, "--show-sdk-path")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get %s SDK path: %w\n%s", sdk, err, string(output))
	}
	return strings.TrimSpace(string(output)), nil
}

func getIOSCC(sdk string) (string, error) {
	cmd := exec.Command("xcrun", "--sdk", sdk, "-f clang")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get %s SDK path: %w\n%s", sdk, err, string(output))
	}
	return string(output), nil
}

func buildIOSArch(arch, sdkPath, cc, sdkName, minOS, output string) error {
	fmt.Printf("Building iOS %s library for %s...\n", arch, sdkName)

	// 设置环境变量
	os.Setenv("GOOS", "darwin")
	os.Setenv("GOARCH", arch)
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", cc)

	// 设置交叉编译标志
	cflags := fmt.Sprintf("-arch %s -isysroot %s -m%s-version-min=%s",
		arch, sdkPath, sdkName, minOS)
	os.Setenv("CGO_CFLAGS", cflags)
	os.Setenv("CGO_LDFLAGS", cflags)

	// 执行构建命令
	cmd := exec.Command("go", "build", "-buildmode=c-archive", "-o", output, ".")
	cmd.Dir = goSrc
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build iOS %s: %w", arch, err)
	}

	fmt.Printf("Build for iOS %s completed successfully\n", arch)
	return nil
}

func createIOSUniversalLibrary(deviceLib, simulatorLib, output string) error {
	fmt.Println("Creating universal library...")

	// 检查 lipo 工具是否可用
	if _, err := exec.LookPath("lipo"); err != nil {
		return fmt.Errorf("lipo tool not found. Make sure Xcode is installed: %w", err)
	}

	cmd := exec.Command("lipo", "-create", deviceLib, simulatorLib, "-output", output)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create universal library: %w", err)
	}

	// 可选：删除中间文件
	os.Remove(deviceLib)
	os.Remove(simulatorLib)

	fmt.Println("Universal library created successfully")
	return nil
}

// BuildIOS compiles the project for iOS.
func BuildIOS() error {
	fmt.Println("Building for iOS...")
	outPath += "ios"
	//arch := os.Getenv("GOARCH")

	// 编译真机版本 (arm64)
	sdkPath, err := getIOSSDKPath("iphoneos")
	if err != nil {
		return err
	}

	cc, err := getIOSCC("iphoneos")
	if err != nil {
		return err
	}

	iphoneosName := soName + "_arm64.a"
	if err := buildIOSArch("arm64", sdkPath, cc, "iphoneos", "13.0", outPath+"/"+iphoneosName); err != nil {
		return err
	}

	// 编译模拟器版本 (x86_64)
	// 获取 Xcode SDK 路径
	//simSdkPath, err := getIOSSDKPath("iphonesimulator")
	//if err != nil {
	//	return err
	//}
	//
	//ccSimulator, err := getIOSCC("iphonesimulator")
	//if err != nil {
	//	return err
	//}

	//iphonesimulatorName := soName + "_x86_64.a"
	//if err := buildIOSArch("x86_64", simSdkPath, ccSimulator, "iphonesimulator", "13.0", outPath+"/"+iphonesimulatorName); err != nil {
	//	return err
	//}

	// 合并为通用库
	iosName := soName + ".a"
	if err := createIOSUniversalLibrary(outPath+"/"+iphoneosName, outPath+"/"+iphonesimulatorName, outPath+"/"+iosName); err != nil {
		return err
	}

	fmt.Println("iOS universal library built successfully at", outPath+"/"+iosName)
	return nil
}

// BuildLinux compiles the project for Linux.
func BuildLinux() error {
	fmt.Println("Building for Linux...")

	outPath += "linux"
	arch := os.Getenv("GOARCH")
	cc := os.Getenv("CC")
	cxx := os.Getenv("CXX")

	if len(arch) == 0 {
		arch = runtime.GOARCH
	}

	if len(cc) == 0 {
		cc = "gcc"
	}

	if len(cxx) != 0 {
		os.Setenv("CXX", cxx)
	}

	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", arch)
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", cc) //

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", outPath+"/"+soName+".so", ".")
	cmd.Dir = goSrc
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to build for Linux: %v\n", err)
		return err
	}
	fmt.Println("Build for Linux completed successfully.")
	return nil
}

// BuildWindows compiles the project for Windows.
func BuildWindows() error {
	fmt.Println("Building for Windows...")

	outPath += "windows"
	arch := os.Getenv("GOARCH")
	cc := os.Getenv("CC")
	cxx := os.Getenv("CXX")

	if len(arch) == 0 {
		arch = runtime.GOARCH
	}

	if len(cc) == 0 {
		cc = "gcc"
	}

	if len(cxx) != 0 {
		os.Setenv("CXX", cxx)
	}

	os.Setenv("GOOS", "windows")
	os.Setenv("GOARCH", arch)
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", cc)

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", outPath+"/"+soName+".dll", ".")
	cmd.Dir = goSrc
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to build for Windows: %v\n", err)
		return err
	}
	fmt.Println("Build for Windows completed successfully.")
	return nil
}
