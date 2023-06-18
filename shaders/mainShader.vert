#version 330
in vec3 vp;

uniform mat4 projection;
uniform mat4 view;
uniform vec2 position;
uniform vec2 scale;

void main() {
    mat4 model = mat4(
            scale.x, 0, 0, 0,
            0, scale.y, 0, 0,
            0, 0, 1, 0,
            position.x, position.y, 0, 1
        );
    gl_Position =  projection * view * model *  vec4(vp, 1.0);
}
