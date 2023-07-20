import 'dart:ffi' as ffi; // For FFI
import 'dart:io' as io;

typedef hello_func = ffi.Void Function(); // FFI fn signature
typedef Hello = void Function(); // Dart fn signature

void main() {
	if (io.Platform.isAndroid) {
		return;
	}

	ffi.DynamicLibrary dylib;
	if (io.Platform.isWindows) {
		dylib = ffi.DynamicLibrary.open('stuff.dll');
	} else {
		dylib = ffi.DynamicLibrary.open('stuff.so');
	}

	sayHello(dylib);
}

void sayHello(ffi.DynamicLibrary dylib) {
	final Hello hello = dylib.lookup<ffi.NativeFunction<hello_func>>('Hello').asFunction();

	hello();
}

