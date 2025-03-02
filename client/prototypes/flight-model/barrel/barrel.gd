extends RigidBody2D

@export var ROTATION_SPEED: float = 2.0

func _physics_process(delta: float) -> void:
	rotation += ROTATION_SPEED * delta  # Rotate right
