export function Connection() {
    return (
        <div className='pb-[30px]'>
            <div className='flex flex-col w-full'>
                <div className=' mr-4 w-full h-[1px] bg-[#2A2929]' />
                <div className='flex flex-row justify-between py-[20px]'>
                    <div className='flex flex-row items-center'>
                        <p className="bg-green-400 w-[16px] h-[16px] rounded-full ml-6" />
                        <p className="text-[#FFFFFF] text-[13px] text-uppercase pl-4 font-light">REDE</p>
                    </div>
                    <p className='text-[#626262] text-[13px] pr-4'>21:04</p>
                </div>
                
                <div className=' mr-4 w-full h-[1px] bg-[#2A2929]' />
                
                <div className='flex flex-row justify-between py-[16px]'>
                    <div className='flex flex-row items-center'>
                        <p className="bg-red-400 w-[16px] h-[16px] rounded-full ml-6" />
                        <p className="text-[#FFFFFF] text-[13px] text-uppercase pl-4 font-light">SERVIDOR</p>
                    </div>
                    <p className='text-[#626262] text-[13px] pr-4'>10:56</p>
                </div>

                <div className=' mr-4 w-full h-[1px] bg-[#2A2929]' />

                <div className='flex flex-row justify-between py-[16px]'>
                    <div className='flex flex-row items-center'>
                        <p className="bg-yellow-400 w-[16px] h-[16px] rounded-full ml-6" />
                        <p className="text-[#FFFFFF] text-[13px] text-uppercase pl-4 font-light">NTP SERVER</p>
                    </div>
                    <p className='text-[#626262] text-[13px] pr-4'>10:56</p>
                </div>
                
            </div>
        </div>
    )
}