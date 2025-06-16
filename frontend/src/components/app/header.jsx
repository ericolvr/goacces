import {  BatteryFull, Menu, Wifi } from 'lucide-react'

export function Header() {
    return (
        <div className="w-full flex px-5 py-4  justify-between items-center bg-[#141414] shadow-xs">
            <div className='flex gap-3'>
                <h4 className='text-white text-[14px] font-semibold'>SMART ACCESS</h4>
            </div> 
            
            <div className="flex gap-4">
                <p className="text-[#777777] text-[13px]">Basa V.1.0.2</p>
            </div>
        </div>
    )
}