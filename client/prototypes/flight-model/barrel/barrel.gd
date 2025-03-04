extends RigidBody2D

# Access the screen size from the GameManager
var spawn_area = GameManager.screen_size

@export var barrel_scene: PackedScene
@export var spawn_count: int = 10

func _ready() -> void:
	print("Parent global_position:", global_position)
	# Spawn the objects
	randomize() #Randomize the seed
	for i in range(spawn_count):
		var instance = barrel_scene.instantiate() as RigidBody2D # Create an instance of the object scene
		var random_rotation = randf_range(0, 360)

		# Random position within the spawn area
		var random_position = Vector2(
			randf_range(-spawn_area.x / 2, spawn_area.x / 2),
			randf_range(-spawn_area.y / 2, spawn_area.y / 2)
		)
		
		add_child(instance)
		instance.position = global_position + random_position
		#instance.position = random_position
		instance.gravity_scale = 0
		instance.rotation_degrees = random_rotation