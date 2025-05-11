extends Area2D

@export var speed: float = 500.0
@export var damage: float = 10.0
@export var lifetime: float = 2.0  # How long the projectile lives before being destroyed

var direction: Vector2 = Vector2.RIGHT  # Default direction, will be set when fired

func _ready():
	# Start the lifetime timer
	await get_tree().create_timer(lifetime).timeout
	queue_free()

func _process(delta):
	# Move the projectile in its direction
	position += direction * speed * delta

func _on_body_entered(body):
	# Check if we hit a barrel
	if body.is_in_group("barrels"):
		# Apply damage to the barrel
		if body.has_method("take_damage"):
			body.take_damage(damage)
		# Destroy the projectile
		queue_free()

func initialize(start_position: Vector2, dir: Vector2):
	position = start_position
	direction = dir.normalized()
	# Rotate the projectile to face its direction
	rotation = direction.angle() 