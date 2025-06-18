import { Header } from '@/components/app/header';
import { Sidebar } from '@/components/app/sidebar';
import { Connection } from '@/components/app/connection';
import { Clock } from '@/components/app/clock';
import { CircleCheckBig, Delete } from 'lucide-react';
import { Keypad } from '@/components/app/keypad';

export function Home() {
  
    return (
        <div className="flex flex-col bg-[#111111] w-full h-[480px]">    
            {/* <Header />
            <div className="flex h-full">
                <Sidebar />
                <div className="bg-[#292929] mt-[40px] w-[1px] h-[325px] " />
                
                <div className='w-[370px] h-screen flex-end justify-end'>    
                    <div className='flex flex-col w-full h-[410px]'>
                        <div className='flex-1 flex justify-center items-center'>
                            <Clock />
                        </div>
                        <Connection />
                    </div>
                </div> */}
                <div className='w-[300px] h-screen'>
                    <Keypad />
                </div>
            </div>    
        // </div>
    );
}
