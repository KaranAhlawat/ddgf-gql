import type { Component } from 'solid-js'
import { useRoutes, useLocation } from 'solid-app-router'

import { routes } from './routes'

const App: Component = () => {
  const location = useLocation()
  const Route = useRoutes(routes)

  return (
    <>
    </>
  )
}

export default App
