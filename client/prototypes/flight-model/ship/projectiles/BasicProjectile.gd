extends Projectile

func _ready():
	speed = 500.0
	damage = 10.0
	lifetime = 2.0
	fire_cooldown = 0.2
	projectile_name = "Basic Projectile"
	super._ready()
	# Start the lifetime timer
	await get_tree().create_timer(lifetime).timeout
	queue_free() 