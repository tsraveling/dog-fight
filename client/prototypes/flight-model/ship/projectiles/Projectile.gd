@tool
extends Area2D
class_name Projectile

# Base properties that child classes will override
var speed: float = 500.0
var damage: float = 10.0
var lifetime: float = 2.0
var fire_cooldown: float = 0.2
var projectile_name: String = "Projectile"
var direction: Vector2 = Vector2.RIGHT

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