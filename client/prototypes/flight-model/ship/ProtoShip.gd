extends RigidBody2D

@export var thrust_force: float = 5000.0
@export var rotation_speed: float = 5.0
@export var torque_force: float = 500.0 # Rotational force

@onready var thrust_polygon := $ThrustPolygon
@onready var projectile_scene := preload("res://prototypes/flight-model/ship/projectiles/LaserProjectile.tscn")

var forward_thrust: Vector2 = Vector2.ZERO
var can_fire: bool = true

func _ready() -> void:
	linear_damp_mode = RigidBody2D.DAMP_MODE_REPLACE
	linear_damp = 0.0
	angular_damp_mode = RigidBody2D.DAMP_MODE_REPLACE
	angular_damp = 0.0

func _physics_process(delta: float) -> void:
	if Input.is_action_pressed("move_up"):
		thrust_polygon.visible = true
		forward_thrust = Vector2(0, -thrust_force * delta) # Apply thrust in the forward direction
		linear_velocity += forward_thrust.rotated(rotation) * delta # Rotates the thruster to point in the direction of the ship
	else:
		thrust_polygon.visible = false

	if Input.is_action_pressed("angular_brake"):
		angular_damp = 4
	else:
		angular_damp = 0.0

	if Input.is_action_pressed("move_left"):
		apply_torque(-torque_force * delta)  # Rotate left
	if Input.is_action_pressed("move_right"):
		apply_torque(torque_force * delta)  # Rotate right
		
	if Input.is_action_pressed("shoot") and can_fire:
		fire_projectile()

func fire_projectile() -> void:
	# Create a new projectile instance
	var projectile = projectile_scene.instantiate()
	get_tree().root.add_child(projectile)
	
	# Calculate spawn position slightly in front of the ship
	var spawn_offset = Vector2(0, -30).rotated(rotation)
	projectile.initialize(global_position + spawn_offset, Vector2(0, -1).rotated(rotation))
	
	# Start cooldown using the projectile's cooldown value
	can_fire = false
	await get_tree().create_timer(projectile.fire_cooldown).timeout
	can_fire = true
