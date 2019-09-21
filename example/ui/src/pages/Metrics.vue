<template>
  <div class="page columns is-centered is-gapless">
    <div class="column is-10-desktop">
      <div class="box">
        <div class="columns">
          <div v-if="loading">Loading...</div>

          <template v-if="data">
            <div v-for="item in data.resources" :key="item.name">
              {{ item.name }}: {{ item.hits }}
            </div>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Metrics',
  data () {
    return {
      loading: true,
      error: null,
      data: null
    }
  },
  created () {
    this.getData()
  },
  methods: {
    getData () {
      this.error = this.metrics = null
      this.loading = true

      return fetch(`/api/ui/metrics`)
        .then(res => res.json())
        .then(data => {
          this.loading = false

          this.data = data
        })
        .catch(err => {
          this.loading = false

          this.error = err.toString()
        })
    }
  }
}
</script>
