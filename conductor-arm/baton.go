components {
  id: "wobbel"
  component: "/conductor-arm/wobbel.script"
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
  properties {
    id: "sensorPosOffset"
    value: "0.0, 150.0, 0.0"
    type: PROPERTY_TYPE_VECTOR3
  }
  properties {
    id: "force"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "damping"
    value: "0.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "wobbelAmount"
    value: "0.001"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/assets/arm.atlas\"\n"
  "default_animation: \"baton\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 67.0
    z: 1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
