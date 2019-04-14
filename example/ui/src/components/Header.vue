<template>
  <div class="header is-flex">
    <div class="header-nav">
      <a class="header-nav-item dropdown is-right is-hoverable">
        <div class="header-nav-icon">
          <svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="3"></circle><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path></svg>
        </div>

        <div class="dropdown-menu" role="menu">
          <div class="dropdown-content">
            <router-link class="dropdown-item" to="/settings/api">
              API Keys
            </router-link>

            <div class="dropdown-divider"></div>

            <a class="dropdown-item" v-on:click="logout">
              Log Out
            </a>
          </div>
        </div>
      </a>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Header',
  methods: {
    logout () {
      this.$store.dispatch('logout')
      this.$router.push('/login')
    }
  }
}
</script>

<style lang="scss">
@import '@/assets/styles/bulma/_utilities.scss';

.header {
  padding: 0 16px;
  background: $white;
  border-bottom: 1px solid rgba(207, 216, 220, .5);
  box-shadow: 0 2px 4px 0 rgba(37, 50, 56, .03);
}

.header-nav {
  display: flex;
  flex: 1 1 auto;
  flex-direction: row;
  justify-content: flex-end;
}

.header-nav-icon {
  width: 20px;
  height: 20px;
  color: $grey-dark;
}

.header-nav-item {
  position: relative;
  display: flex;
  align-items: center;
  height: 55px;
  padding: 13px 1em;

  > span {
    font-size: $size-normal;
  }

  &::after {
    content: "";
    position: absolute;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 3px;
    background: transparent;
  }
}

.header-nav-item:hover,
.header-nav-item:focus {
  &::after {
    background: $grey-dark;
  }
}

.header-nav-item.is-active {
  &::after {
    background: $corporate-green;
  }
}

.header .dropdown-menu {
  padding-top: 1px;

  > .dropdown-content {
    border-top-left-radius: 0;
    border-top-right-radius: 0;
  }
}
</style>
