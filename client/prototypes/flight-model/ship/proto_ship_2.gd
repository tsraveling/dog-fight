extends RigidBody2D

@export var THRUST_FORCE: float = 5000.0
@export var ROTATION_SPEED: float = 5.0

var forward_thrust: Vector2 = Vector2.ZERO

func _physics_process(delta: float) -> void:
	if Input.is_action_pressed("move_up"):
		# velocity += Vector2(0, -THRUST_FORCE * delta)  # Apply thrust
		forward_thrust = Vector2(0, -THRUST_FORCE * delta) # Apply thrust in the forward direction
		linear_velocity += forward_thrust.rotated(rotation) * delta # Rotates the thruster to point in the direction of the ship

	if Input.is_action_pressed("move_left"):
		rotation -= ROTATION_SPEED * delta  # Rotate left
	if Input.is_action_pressed("move_right"):
		rotation += ROTATION_SPEED * delta  # Rotate right


