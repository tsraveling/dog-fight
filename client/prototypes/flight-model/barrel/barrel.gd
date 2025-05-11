extends RigidBody2D

@export var max_health: float = 20.0
@onready var anim = $ExplosionAnimation
@onready var barrel = $Sprite2D

var current_health: float

func _ready() -> void:
	linear_damp_mode = RigidBody2D.DAMP_MODE_REPLACE # Not working when located in parent
	linear_damp = 0.0
	angular_damp_mode = RigidBody2D.DAMP_MODE_REPLACE
	angular_damp = 0.0

	current_health = max_health
	anim.animation_finished.connect(_on_animation_finished) # Attach signal for end of animation
	add_to_group("barrels")  # Add to barrels group for projectile detection

func _on_body_entered(body):
	if body.name == "ProtoShip":
		trigger_explosion() # On collision with ship, trigger the explosion

func take_damage(amount: float) -> void:
	current_health -= amount
	if current_health <= 0:
		trigger_explosion()

func trigger_explosion():
	anim.visible = true # Keeps the animation hidden until triggered
	anim.play("explosion_animation")
	barrel.visible = false

func _on_animation_finished():
	queue_free() # Free the instance after the animation is finished
