const basePath = `/api`

export function request (path) {
  return fetch(`${basePath}/${path}`)
}

export function fetchMeta () {
  return fetch(`${basePath}/ui/meta`)
    .then(res => res.json())
}

export function auth (email, password) {
  const requestOpts = {
    method: 'POST',
    body: JSON.stringify({ email, password })
  }

  return fetch(`${basePath}/auth`, requestOpts)
    .then(res => res.json())
}

export default request
