#include "GLFW/glfw3.h"

void goCursorPosCallback(GLFWwindow *window, double x, double y);
void goSizeCallback(GLFWwindow *window, int width, int height);
void goFramebufferSizeCallback(GLFWwindow *window, int width, int height);
void goKeyCallback(GLFWwindow *window, int key, int scancode, int action, int mods);
void goMouseButtonCallback(GLFWwindow *window, int button, int action, int mods);
void goScrollCallback(GLFWwindow *window, double xoff, double yoff);

void glfwSetCursorPosCallback_fix(GLFWwindow *window) {
    glfwSetCursorPosCallback(window, (GLFWcursorposfun)goCursorPosCallback);
}

void glfwSetWindowSizeCallback_fix(GLFWwindow *window) {
    glfwSetWindowSizeCallback(window, (GLFWwindowsizefun)goSizeCallback);
}

void glfwSetFramebufferSizeCallback_fix(GLFWwindow *window) {
    glfwSetFramebufferSizeCallback(window, (GLFWframebuffersizefun)goFramebufferSizeCallback);
}

void glfwSetKeyCallback_fix(GLFWwindow *window) {
    glfwSetKeyCallback(window, (GLFWkeyfun)goKeyCallback);
}

void glfwSetMouseButtonCallback_fix(GLFWwindow *window) {
    glfwSetMouseButtonCallback(window, (GLFWmousebuttonfun)goMouseButtonCallback);
}

void glfwSetScrollCallback_fix(GLFWwindow *window) {
    glfwSetScrollCallback(window, (GLFWscrollfun)goScrollCallback);
}