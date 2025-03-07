
extends RigidBody2D

@onready var anim = $ExplosionAnimation
@onready var barrel = $Sprite2D

func _ready() -> void:
	linear_damp_mode = RigidBody2D.DAMP_MODE_REPLACE # Not working when located in parent
	linear_damp = 0.0
	angular_damp_mode = RigidBody2D.DAMP_MODE_REPLACE
	angular_damp = 0.0

	anim.animation_finished.connect(_on_animation_finished) # Attach signal for end of animation
	# connect("body_entered", _on_body_entered)  # Attach signal for collision detection

func _on_body_entered(body):
	if body.name == "ProtoShip":
		trigger_explosion() # On collision, trigger the explosion

func trigger_explosion():
	anim.visible = true # Keeps the animation hidden until triggered
	anim.play("explosion_animation")
	barrel.visible = false

func _on_animation_finished():
	queue_free() # Free the instance after the animation is finished
