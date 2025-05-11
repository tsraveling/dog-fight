extends Projectile

func _ready():
	# Override properties for laser
	speed = 800.0  # Faster than basic projectile
	damage = 50.0  # More damage
	lifetime = 10.0  # Longer lifetime
	fire_cooldown = 0.8  # Slightly longer cooldown
	projectile_name = "Laser"
	super._ready()
	
	# Start the lifetime timer
	await get_tree().create_timer(lifetime).timeout
	queue_free()

func _on_body_entered(body):
	# Check if we hit a barrel
	if body.is_in_group("barrels"):
		# Apply damage to the barrel
		if body.has_method("take_damage"):
			body.take_damage(damage)
		# Don't destroy the projectile - it's a piercing laser 