<template>
  <div>
    <Spinner v-if="!items && !error" />

    <div v-if="error">
      {{ error }}
    </div>

    <div class="data-table" :class="{ 'is-loading': loading }">
      <table class="table is-fullwidth" v-if="items">
        <thead>
          <tr>
            <th v-for="(_, key) in items[0]" :key="key">{{ key }}</th>
            <th class="data-table-actions"></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in items" :key="item.id">
            <td v-for="(val, _, index) in item" :key="item.id + '-' + index"> {{ val }}</td>
            <td class="data-table-actions">
              <button class="button is-link">
                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="16 3 21 8 8 21 3 21 3 16 16 3"></polygon></svg>
              </button>
              <button class="button is-danger is-outlined">
                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path><line x1="10" y1="11" x2="10" y2="17"></line><line x1="14" y1="11" x2="14" y2="17"></line></svg>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <nav class="pagination" role="navigation" aria-label="pagination" v-if="paginate">
      <a class="pagination-previous" @click="filterItems({ offset: filters.offset - filters.limit })" :disabled="filters.offset == 0">Previous</a>
      <a class="pagination-next" @click="filterItems({ offset: filters.offset + filters.limit })">Next page</a>
      <ul class="pagination-list">
        <li v-for="n in 5" :key="n">
          <a class="pagination-link" :class="{ 'is-current': filters.offset === (n - 1) * filters.limit}" aria-label="Goto page 1" @click="filterItems({ offset: (n - 1) * filters.limit })">{{ n }}</a>
        </li>
      </ul>
    </nav>
  </div>
</template>

<script>
import Spinner from '@/components/Spinner'

export default {
  name: 'DataTable',
  props: ['uri', 'paginate'],
  components: {
    Spinner
  },
  data () {
    return {
      loading: true,
      error: null,
      filters: {
        offset: 0,
        limit: 10
      },
      items: null
    }
  },
  created () {
    this.getItems()
  },
  watch: {
    // call again the method if the route changes
    // https://router.vuejs.org/guide/advanced/data-fetching.html#fetching-after-navigation
    '$route': 'getItems'
  },
  methods: {
    getItems () {
      this.error = null
      this.loading = true

      let query = ''
      if (this.filters.offset !== 0) {
        query += `?offset=${this.filters.offset}`
      }

      return fetch(`/api/${this.uri}${query}`)
        .then(res => res.json())
        .then(data => {
          this.items = data
        })
        .catch(err => {
          this.error = err.toString()
        })
        .finally(_ => {
          this.loading = false
        })
    },

    filterItems (parameters) {
      if (parameters.hasOwnProperty('offset') && parameters.offset < 0) {
        parameters.offset = 0
      }

      this.filters = Object.assign(this.filters, parameters)
      this.getItems()
    }
  }
}
</script>

<style lang="scss">
@import '@/assets/styles/bulma/_utilities.scss';

.data-table {
  td {
    vertical-align: middle;
  }
}

.data-table {
  &-actions {
    background: $white-bis;
    text-align: center;
    max-width: 120px;
  }
}

.data-table + .pagination {
  margin-top: .7rem;
}

.data-table.is-loading {
  opacity: 0.3;
  cursor: progress;
}
</style>
