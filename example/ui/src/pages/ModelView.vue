<template>
  <Page class="columns is-centered is-gapless">
    <div class="column is-10-desktop">
      <div class="box">
        <DataTable :uri="$route.params.name" :paginate="true" />
      </div>
    </div>
  </Page>
</template>

<script>
import Page from '@/components/Page'
import DataTable from '@/components/DataTable'

export default {
  name: 'ModelView',
  components: {
    Page,
    DataTable
  }
}
</script>
