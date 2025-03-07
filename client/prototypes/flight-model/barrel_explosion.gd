extends AnimatedSprite2D

@onready var sprite = $AnimatedSprite2D

func _ready() -> void:
	sprite.play("explosion")
	await sprite.animation_finished()
	queue_free()
 
 
