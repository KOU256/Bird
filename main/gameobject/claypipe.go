components {
  id: "claypipe"
  component: "/main/scripts/claypipe.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "bottom_claypipe"
  type: "sprite"
  data: "tile_set: \"/main/atlases/claypipe.atlas\"\n"
  "default_animation: \"claypipe\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: -224.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "top_claypipe"
  type: "sprite"
  data: "tile_set: \"/main/atlases/claypipe.atlas\"\n"
  "default_animation: \"claypipe\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 224.0
    z: 0.0
  }
  rotation {
    x: 1.0
    y: 0.0
    z: 0.0
    w: 6.123234E-17
  }
}
