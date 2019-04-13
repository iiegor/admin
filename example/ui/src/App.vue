<template>
  <div id="app">
    <div class="env-banner level development" v-if="env == 'development'">
      <div class="level-item notification has-background-warning has-text-grey-dark">Local Development</div>
    </div>

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
  </div>
</template>

<script>
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
  mounted () {
    this.fetchMeta()
  },
  methods: {
    fetchMeta () {
      fetch('/api/ui/meta')
        .then(res => res.json())
        .then(data => {
          this.$store.commit('meta', data)

          this.loading = false
        })
        .catch(err => {
          console.error('fetchMeta', err)

          // retry meta fetch after 5s
          setTimeout(() => {
            this.fetchMeta()
          }, 5000)
        })
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
