import './App.css'
import { Navbar } from './components/navbar/navbar'
import { Form } from './components/form/Form'

function App() {
  return (
    <>
      <div>
        <Navbar />
        <div className='text-center'>
          <br />
          <h1 className="bold text-3xl">Submit your Github URL!!</h1>
        </div>
          <Form />
      </div>
    </>
  )
}

export default App
