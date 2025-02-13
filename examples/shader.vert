#version 460

layout(location = 0) in vec3 aPosition;
layout(location = 1) in vec3 aColor;
layout(location = 2) in vec2 aUv;
layout(location = 3) in vec3 aNormal;

out vec3 vColor;
out vec2 vUv;
out vec3 vNormal;

uniform mat4 matrix;

void main() {
  gl_Position = matrix * vec4(aPosition, 1.0);
  vColor = aColor;
  vUv = aUv;
  vNormal = aNormal;
}
