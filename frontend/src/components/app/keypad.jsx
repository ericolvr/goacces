import { CircleCheckBig, Delete } from 'lucide-react';

export function Keypad() {
  return (
    <div className="flex w-[320px] flex-col justify-start">
      {/* input */}
      <div className='px-14 h-[30px] mt-3'>
        <div className='h-[30px]'>
          <input type="text" className="w-full h-full bg-white"  />
        </div>
      </div>
      {/* keypad */}
      <div className='flex items-center mt-7 ml-10'>
        <div className="grid grid-cols-3 gap-3 w-[228px]">
          <button className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]">1</button>
          <button className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]">2</button>
          <button className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]">3</button>

          <button className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]">4</button>
          <button className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]">5</button>
          <button className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]">6</button>

          <button className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]">7</button>
          <button className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]">8</button>
          <button className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]">9</button>

          <button className="w-[68px] h-[68px] rounded-md bg-red-400 text-white text-2xl font-regular shadow hover:bg-[#3a1a1a] flex items-center justify-center">
            <Delete className="size-6" />
          </button>
          <button className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]">0</button>
          <button className="w-[68px] h-[68px] rounded-md bg-green-400 text-white text-2xl font-regular shadow hover:bg-green-680 flex items-center justify-center">
            <CircleCheckBig className="size-6" />
          </button>
        </div>
      </div>
    </div>
  );
}
