<template>
  <div class="columns is-centered is-gapless">
    <div class="column is-10-desktop">
      <div class="box">
        <form action="" method="POST" v-on:submit.prevent="onSubmit">
          <div class="field">
            <label for="email" class="label">Email</label>
            <div class="control">
              <input type="email" class="input" id="email" name="email" v-model="email">
            </div>
          </div>

          <div class="field">
            <label for="password" class="label">Password</label>
            <div class="control">
              <input type="password" class="input" id="password" name="password" v-model="password">
            </div>
          </div>

          <div class="field">
            <div class="control">
              <button type="submit" class="button is-link">Log in</button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    onSubmit (evt) {
      this.$store.dispatch('login', {
        email: this.email,
        password: this.password
      })
    }
  }
}
</script>
