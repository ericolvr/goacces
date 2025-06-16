import { 
    CalendarClock, 
    Cctv, 
    Menu, 
    Settings2, Users 

} from 'lucide-react';

export function Sidebar() {
    return (
        <div className="w-[120px] h-full">
            <div className="px-8 py-10">
                <div className="flex flex-col gap-8">
                    <div className="w-[54px] h-[54px] bg-[#212121] opacity-50 rounded-full flex justify-center items-center hover:opacity-100">
                        <Settings2 color="#FFFFFF" size={22} />
                    </div>
                    <div className="w-[54px] h-[54px] bg-[#212121] opacity-50 rounded-full flex justify-center items-center hover:opacity-100">
                    <CalendarClock color="#FFFFFF" size={23} />
                    </div>
                    <div className="w-[54px] h-[54px] bg-[#212121] opacity-50 rounded-full flex justify-center items-center hover:opacity-100">
                        <Cctv color="#FFFFFF" size={23} />
                    </div>
                    <div className="w-[54px] h-[54px] bg-[#212121] opacity-50 rounded-full flex justify-center items-center hover:opacity-100">
                        <Users color="#FFFFFF" size={23} />
                    </div>
                </div>
            </div>
        </div>
    )
}