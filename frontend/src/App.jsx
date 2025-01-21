import './App.css'
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import LandingPage from './pages/landingPage/landingPage'
import RegisterPage from './pages/registerPage/registerPage'
import SignInPage from './pages/signInPage/signInPage'

function App() {

  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route index element={<LandingPage />} />
          <Route path="/register" element={<RegisterPage/>} />
          <Route path="/sign-in"element={<SignInPage/>} />
          </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
