package glfw

//#include<GL/glfw.h>
/*
void otherCallback(void (*f)(int x,int y)){
     f(1,2);
}

void SetupCallbacks(){
 extern void GLFWCALL mouseButtonCallback(int,int);
 extern void GLFWCALL keyCallback(int key, int action);
 extern void GLFWCALL mousePosCallback(int x, int y);
 extern int GLFWCALL windowCloseCallback();
 extern void GLFWCALL windowResizeCallback(int w, int h);
 extern void GLFWCALL mouseWheelCallback(int pos);

 glfwSetMouseButtonCallback(mouseButtonCallback);
 glfwSetMousePosCallback(mousePosCallback);
 glfwSetMouseWheelCallback(mouseWheelCallback);
 glfwSetKeyCallback(keyCallback);
 glfwSetWindowSizeCallback(windowResizeCallback);
 glfwSetWindowCloseCallback(windowCloseCallback);
}

*/
import "C"
import "fmt"
import "container/list"

type _GLFWEventCore struct{
	Listeners *list.List
}

var eventCore *_GLFWEventCore


//import "unsafe"
type IKeyEvent interface{
     KeyUp(key int)
     KeyDown(key int)
}

type IMouseEvent interface {
     Move(x int, y int);
      ClickDown(button int)
     ClickUp(button int)
     Drag(x int, y int, dx float32, dy float32)
}

type KeyEvent struct{
	Key int
	Action int
}

type MouseClickEvent struct{
	Button int
	Action int
}

type MouseMoveEvent struct{
	X int
	Y int
}

type MouseWheelEvent struct{
	Pos int
}


type WindowEvent struct{
	closed bool
}

type WindowResizeEvent struct{
	X int
	Y int
}

func Init(width int, height int){
     eventCore = new(_GLFWEventCore)
     eventCore.Listeners = new(list.List)
	C.glfwInit();
     C.glfwOpenWindow(_Ctype_int(width),_Ctype_int(height),8,8,8,8,8,8,C.GLFW_WINDOW);
     C.SetupCallbacks()
	
	
}   
func AddListener(listener interface{}){
	eventCore.AddListener(listener)
}

func (evc *_GLFWEventCore) AddListener(a interface{}){
	evc.Listeners.PushBack(a)
}


func listmap(f func(interface{}),lst *list.List){
	for e:=lst.Front(); e != nil; e = e.Next() {
		f(e.Value)
	}
}

//export keyCallback
func keyCallback(key int, action int){
	for e := eventCore.Listeners.Front(); e != nil; e = e.Next() {
		fent, ok :=e.Value.(func(KeyEvent))
		if ok {
			fent(KeyEvent{key,action})
		}
	}
}

//export mouseButtonCallback
func mouseButtonCallback(button int, action int){
     for e := eventCore.Listeners.Front(); e != nil; e = e.Next() {
		fent,ok := e.Value.(func(MouseClickEvent))
		if ok {
			fent(MouseClickEvent{button,action})
		}
	}
}

//export mousePosCallback
func mousePosCallback(x int, y int){
	listmap(func(i interface{}){
		fent,ok := i.( func(MouseMoveEvent) )
		if ok {
			fent(MouseMoveEvent{x,y})
		}
	}, eventCore.Listeners)
}
//export mouseWheelCallback
func mouseWheelCallback(pos int){
	listmap(func(i interface{}){
		fent,ok := i.( func(MouseWheelEvent) )
		if ok {
			fent(MouseWheelEvent{pos})
		}
	}, eventCore.Listeners)
}

//export windowResizeCallback
func windowResizeCallback(width int, height int){
	listmap(func(i interface{}){
		fent,ok := i.( func(WindowResizeEvent) )
		if ok {
			fent(WindowResizeEvent{width,height})
		}
	}, eventCore.Listeners)
}

//export windowCloseCallback
func windowCloseCallback() _Ctype_int{
     return C.GL_TRUE;
}

func SwapBuffers(){
     C.glfwSwapBuffers()
}

func Version(){
	fmt.Println("v1")
}

const (
	KEY_SPACE = 32
	KEY_A = 65
	KEY_B = 66
	KEY_C = 67
	KEY_D = 68
	KEY_E = 69
	KEY_F = 70
	KEY_G = 71
	KEY_H = 72
	KEY_I = 73
	KEY_J = 74
	KEY_K = 75
	KEY_L = 76
	KEY_M = 77
	KEY_N = 78
	KEY_O = 79
	KEY_P = 80
	KEY_Q = 81
	KEY_R = 82
	KEY_S = 83
	KEY_T = 84
	KEY_U = 85
	KEY_V = 86
	KEY_W = 87
	KEY_X = 88
	KEY_Y = 89
	KEY_Z = 90
	KEY_ESC = 257
	KEY_F1 = 258
	KEY_F2 = 259
	KEY_F3 = 260
	KEY_F4 = 261
	KEY_F5 = 262
	KEY_F6 = 263
	KEY_F7 = 264
	KEY_F8 = 265
	KEY_F9 = 266
	KEY_F10 = 267
	KEY_F11 = 268
	KEY_F12 = 269
	KEY_UP = 283
	KEY_DOWN = 284
	KEY_LEFT = 285
	KEY_RIGHT = 286
	KEY_LSHIFT = 287
	KEY_RSHIFT = 288
	KEY_LCTRL = 289
	KEY_RCTRL = 290
	KEY_ALT = 291
	
	KEY_ENTER = 294
)
