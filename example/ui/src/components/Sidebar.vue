<template>
  <div class="sidebar is-fullheight">
    <router-link class="sidebar-item" active-class="is-active" to="/" exact>
      <div class="sidebar-icon">
        <svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path><polyline points="9 22 9 12 15 12 15 22"></polyline></svg>
      </div>

      <h2>Dashboard</h2>
    </router-link>

    <router-link v-for="item in $store.state.meta.resources" :key="item.endpoint" class="sidebar-item" active-class="is-active" :to="'/resource/' + item.endpoint" exact>
      <div class="sidebar-icon">
        <svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="8" y1="6" x2="21" y2="6"></line><line x1="8" y1="12" x2="21" y2="12"></line><line x1="8" y1="18" x2="21" y2="18"></line><line x1="3" y1="6" x2="3" y2="6"></line><line x1="3" y1="12" x2="3" y2="12"></line><line x1="3" y1="18" x2="3" y2="18"></line></svg>
      </div>

      <h2>{{ item.label }}</h2>
    </router-link>

    <router-link class="sidebar-item" active-class="is-active" to="/metrics">
      <div class="sidebar-icon">
        <svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="20" x2="18" y2="10"></line><line x1="12" y1="20" x2="12" y2="4"></line><line x1="6" y1="20" x2="6" y2="14"></line></svg>
      </div>

      <h2>Metrics</h2>
    </router-link>
  </div>
</template>

<script>
export default {
  name: 'Navbar'
}
</script>

<style lang="scss">
@import '@/assets/styles/bulma/_utilities.scss';

.sidebar {
  width: 90px;
  padding-top: 16px;
  background: #263238;
  border-right: 1px solid rgba(0, 0, 0, .1);
}

.sidebar-item {
  display: block;
  text-align: center;
  transition: opacity 150ms ease;
  padding: 12px 0;

  > h2 {
    font-size: $size-small;
    color: rgba($white, 0.6);
    text-transform: uppercase;
    letter-spacing: 0.5px;
    padding: 0 5px;
    word-break: break-word;
    white-space: normal;
  }
}

.sidebar-item.is-active,
.sidebar-item:hover {
  .sidebar-icon {
    color: rgba($white, 1);
    background-color: rgba(54, 71, 79, 1);
    transition:
      color 150ms ease,
      background-color 150ms ease;
  }

  > h2 {
    color: rgba($white, 1);
    transition: color 150ms ease;
  }
}

.sidebar-icon {
  position: relative;
  width: 38px;
  height: 38px;
  margin: 0 auto;
  margin-bottom: 8px;
  color: rgba($white, 0.6);
  background-color: rgba(54, 71, 79, .5);
  border-radius: 50%;

  svg {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 20px;
    height: 20px;
    transform: translateY(-50%) translateX(-50%);
  }

  .sidebar-item.is-active & {
    background-color: $link;
  }
}
</style>
