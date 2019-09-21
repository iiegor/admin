<template>
  <div id="app">
    <div class="env-banner level development" v-if="env == 'development'">
      <div class="level-item notification has-background-warning has-text-grey-dark">Local Development</div>
    </div>

    <template v-if="$store.state.loggedIn">
      <div class="columns is-gapless is-fullwidth is-fullheight">
        <div class="column is-narrow" v-if="!loading">
          <Sidebar />
        </div>

        <div class="column view">
          <Spinner v-if="loading" />

          <Header v-if="!loading" />
          <router-view class="view-container is-fullheight" v-if="!loading" />
        </div>
      </div>
    </template>

    <template v-if="!$store.state.loggedIn">
      <router-view class="view-container is-fullheight" v-if="!loading" />
    </template>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import Sidebar from '@/components/Sidebar'
import Header from '@/components/Header'
import Spinner from '@/components/Spinner'

export default {
  name: 'App',
  components: {
    Sidebar,
    Header,
    Spinner
  },
  data () {
    return {
      env: process.env.NODE_ENV,
      loading: true
    }
  },
  computed: {
    ...mapState([
      'loggedIn',
      'meta'
    ])
  },
  mounted () {
    this.fetchMeta()
  },
  watch: {
    loggedIn (val, oldVal) {
      // Viene de hacer login
      // por lo que le redirijimos
      if (val) {
        this.$router.replace('/')
      }

      this.fetchMeta()
    },

    meta (val, oldVal) {
      this.loading = false
    }
  },
  methods: {
    fetchMeta () {
      if (this.loggedIn) {
        this.loading = true

        this.$store.dispatch('fetchMeta')
      } else {
        this.loading = false
      }
    }
  }
}
</script>

<style lang="scss">
@import '@/assets/styles/bulma/bulma.scss';
@import '../static/fonts/inter-ui.css';

html, body, #app {
  height: 100%;
  width: 100%;
  overflow: hidden;
}

.env-banner {
  margin-bottom: 0 !important;
}

.env-banner > .notification {
  position: relative;
  padding-top: 0.3em;
  padding-bottom: 0.3em;
  border-radius: 0;

  &::before {
    content: "";
    position: absolute;
    bottom: -5px;
    left: 40px;
    width: 10px;
    height: 10px;
    background: $yellow;
    transform: rotate(45deg);
  }
}

.view {
  background: rgba(236, 239, 241, .6);
}

.view-container {
  overflow-y: auto;
}
</style>
