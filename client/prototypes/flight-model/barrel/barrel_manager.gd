extends Node2D

@export var barrel_scene: PackedScene
@export var spawn_count: int = 10

# Stores the width (screen_size.x) and height (screen_size.y) of the game window
var spawn_area: Vector2 = Vector2(1920, 1080)

func _ready() -> void:
	if not barrel_scene:
		print("Assign barrel.tscn to the Barrel Scene in the Inspector.")
		return
	
	# Spawn the objects
	for i in range(spawn_count):
		var instance = barrel_scene.instantiate() as RigidBody2D # Create an instance of the object scene
		var random_rotation = randf_range(0, 360) # Random starting rotation
		var random_impulse = Vector2(randf_range(-100, 100), randf_range(-100, 100)) # Random impulse for movement

		# Random position within the spawn area
		var random_position = Vector2(
			randf_range(-spawn_area.x / 2, spawn_area.x / 2),
			randf_range(-spawn_area.y / 2, spawn_area.y / 2)
		)

		instance.position = global_position + random_position # Set the position of the instance
		instance.gravity_scale = 0 # Turns off gravity
		instance.linear_damp = 0 # Not working yet
		instance.angular_damp = 0 # Not working yet
		instance.rotation_degrees = random_rotation # Random starting rotation
		instance.apply_impulse(random_impulse, Vector2.ZERO) # Applies the impulse to the instance
		add_child(instance)
