import tensorflow as tf

# Model Parameters
W = tf.Variable([-1.0], tf.float32)
b = tf.Variable([1.0], tf.float32)

# Inputs and Outputs
x = tf.placeholder(tf.float32)

linear_model = W * x + b

y = tf.placeholder(tf.float32)

# Loss
sqaured_delta = tf.square(linear_model - y)
loss = tf.reduce_sum(sqaured_delta)

init = tf.global_variables_initializer()

sess = tf.Session()
sess.run(init)


print(sess.run(loss, {
  x:[1,2,3,4],
  y:[0,-1,-2,-3],
}))