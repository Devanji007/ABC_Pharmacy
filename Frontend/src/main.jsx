import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'
import { Route, RouterProvider, createBrowserRouter, createRoutesFromElements } from 'react-router-dom'
import HomePage from './pages/homePage.jsx'
import InvoicePage from './pages/invoicePage.jsx'
import ItemPage from './pages/itemPage.jsx'
import CreateItemPage from './pages/createItemPage.jsx'
import CreateInvoicePage from './pages/createInvoicePage.jsx'

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path='/' element={ <App /> }>
      
    <Route index={ true } path='/' element={ <HomePage /> } />
     <Route path='/invoice' element={ <InvoicePage /> } /> 
     <Route path='/item' element={ <ItemPage /> } /> 
     <Route path='/AddItem' element={ <CreateItemPage /> } /> 
     <Route path='/CreateInvoice' element={ <CreateInvoicePage /> } /> 



  </Route>
  )
)


ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={ router } />
  </React.StrictMode>,
)
