#version 330
out vec4 fragColor;

in vec2 textCoord;
flat in int useColor;

uniform vec3 color;
uniform sampler2D text;

void main() {
    if (useColor != 0){
        fragColor = vec4(color, 1.0);
    } else {
        fragColor = texture(text, textCoord);
    }
}
