components {
  id: "bubble"
  component: "/main/bubble.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"bubble\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/bubbles.atlas\"\n"
  "}\n"
  ""
}
