extends CharacterBody2D

const SPEED = 300.0
@export var THRUST_FORCE: float = 200.0

func _physics_process(delta: float) -> void:
	if Input.is_action_pressed("move_up"):
		velocity += Vector2(0, -THRUST_FORCE * delta)  # Apply thrust
		
	move_and_slide()



