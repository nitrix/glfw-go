CMAKE_MINIMUM_REQUIRED(VERSION 3.16)
PROJECT(glfw-go)

SET(CMAKE_CXX_STANDARD 17)

SET(GLFW_BUILD_X11 ON)
SET(GLFW_BUILD_WAYLAND OFF)
SET(GLFW_INSTALL OFF)

SET(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -O3")
SET(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -O3")

ADD_SUBDIRECTORY(thirdparty/glfw)