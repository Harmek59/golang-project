#version 330
layout (location=0) in vec3 vp;
layout (location=1) in vec2 tc;
layout (location=2) in vec2 position;
layout (location=3) in vec2 scale;
layout (location=4) in vec2 textureBegin;
layout (location=5) in vec2 textureEnd;
layout (location=6) in float textID;

out vec2 textCoord;
flat out int textureID;

uniform mat4 projection;
uniform mat4 view;

void main() {
    textCoord = tc * (textureEnd - textureBegin) + textureBegin;
    textureID = int(textID);
    mat4 model = mat4(
            scale.x, 0, 0, 0,
            0, scale.y, 0, 0,
            0, 0, 1, 0,
            position.x + scale.x/2, position.y + scale.y/2, 0, 1
        );
    gl_Position =  projection * view * model *  vec4(vp, 1.0);
}
