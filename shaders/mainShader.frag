#version 330
out vec4 fragColor;

in vec2 textCoord;
flat in int textureID;

uniform vec3 color;
uniform sampler2D texture0;
uniform sampler2D texture1;
uniform sampler2D texture2;
uniform sampler2D texture3;
uniform sampler2D texture4;
uniform sampler2D texture5;
uniform sampler2D texture6;
uniform sampler2D texture7;

void main() {
    if (textureID == -1) {
        fragColor = vec4(color, 1.0);
    } else if (textureID == 0) {
        fragColor = texture(texture0, textCoord);
    } else if (textureID == 1) {
        fragColor = texture(texture1, textCoord);
    } else if (textureID == 2) {
        fragColor = texture(texture2, textCoord);
    } else if (textureID == 3) {
        fragColor = texture(texture3, textCoord);
    } else if (textureID == 4) {
        fragColor = texture(texture4, textCoord);
    } else if (textureID == 5) {
        fragColor = texture(texture5, textCoord);
    } else if (textureID == 6) {
        fragColor = texture(texture6, textCoord);
    } else if (textureID == 7) {
        fragColor = texture(texture7, textCoord);
    } else {
        fragColor = vec4(1.0, 0.0, 0.0, 1.0);
    }
}
