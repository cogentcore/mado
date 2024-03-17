package glfw

import (
	"fmt"
	"unsafe"
)

func (w *Window) initContext() {
	if ctx, err := w.callbacks.D.NewContext(); err != nil {
		panic(err)
	} else {
		w.ctx = ctx
		theApp.Ctx = ctx
	}
	panicError()
}

func (w *Window) MakeContextCurrent() {
	if err := w.ctx.Lock(); err != nil {
		panic(err)
	}
	panicError()
}

// DetachCurrentContext detaches the current context.
func DetachCurrentContext() {
	theApp.Ctx.Unlock()
	panicError()
}

// GetCurrentContext returns the window whose context is current.
func GetCurrentContext() *Window {
	return theApp.MainWindow
}

// SwapBuffers swaps the front and back buffers of the window. If the
// swap interval is greater than zero, the GPU driver waits the specified number
// of screen updates before swapping the buffers.
func (w *Window) SwapBuffers() {
	if err := w.ctx.SwapBuffers(); err != nil {
		panic(err)
	}
	panicError()
}

// SwapInterval sets the swap interval for the current context, i.e. the number
// of screen updates to wait before swapping the buffers of a window and
// returning from SwapBuffers. This is sometimes called
// 'vertical synchronization', 'vertical retrace synchronization' or 'vsync'.
//
// Contexts that support either of the WGL_EXT_swap_control_tear and
// GLX_EXT_swap_control_tear extensions also accept negative swap intervals,
// which allow the driver to swap even if a frame arrives a little bit late.
// You can check for the presence of these extensions using
// ExtensionSupported. For more information about swap tearing,
// see the extension specifications.
//
// Some GPU drivers do not honor the requested swap interval, either because of
// user settings that override the request or due to bugs in the driver.
func SwapInterval(interval int) {
	theApp.Ctx.SwapInterval(interval)
	panicError()
}

// ExtensionSupported reports whether the specified OpenGL or context creation
// API extension is supported by the current context. For example, on Windows
// both the OpenGL and WGL extension strings are checked.
//
// As this functions searches one or more extension strings on each call, it is
// recommended that you cache its results if it's going to be used frequently.
// The extension strings will not change during the lifetime of a context, so
// there is no danger in doing this.
func ExtensionSupported(extension string) bool {
	fmt.Println("not implemented")
	panicError()
	return false
}

// GetProcAddress returns the address of the specified OpenGL or OpenGL ES core
// or extension function, if it is supported by the current context.
//
// A context must be current on the calling thread. Calling this function
// without a current context will cause a GLFW_NO_CURRENT_CONTEXT error.
//
// This function is used to provide GL proc resolving capabilities to an
// external C library.
func GetProcAddress(procname string) unsafe.Pointer {
	fmt.Println("not implemented")
	panicError()
	return nil
}
