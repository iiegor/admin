<template>
  <div class="columns is-centered is-gapless">
    <div class="column is-4 is-offset-4">
      <div class="box">
        <form action="" method="POST" v-on:submit.prevent="onSubmit">
          <div class="field">
            <label for="username" class="label">Username</label>
            <div class="control">
              <input type="text" class="input" id="username" name="username" v-model="username">
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
  name: 'Auth',
  data () {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    onSubmit (evt) {
      this.$store.dispatch('auth', {
        username: this.username,
        password: this.password
      })
    }
  }
}
</script>
