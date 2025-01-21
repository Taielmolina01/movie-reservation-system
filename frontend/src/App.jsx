import './App.css'
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import LandingPage from './pages/landingPage'

function App() {

  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route index element={<LandingPage />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
