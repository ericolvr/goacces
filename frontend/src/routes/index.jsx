import { 
    BrowserRouter, 
    Route, 
    Routes
} from 'react-router-dom'

import { Home } from '@/pages/home'
import { Face } from '@/pages/face'


export const AppRoutes = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/face" element={<Face />} />
            </Routes>
        </BrowserRouter>
    )
}

