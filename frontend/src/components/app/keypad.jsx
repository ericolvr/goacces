import { useEffect, useState } from 'react';
import { CircleCheckBig, Delete } from 'lucide-react';
import { useNavigate } from 'react-router-dom';

import { PinExists } from '../../../wailsjs/go/main/App';


export function Keypad() {
	const nav = useNavigate();
	const [value, setValue] = useState("");
	const [document, setDocument] = useState("");

	const handleClick = (num) => {
		setValue((prev) => (prev.length < 7 ? prev + num : prev));
	};

	const handleDelete = () => {
		setValue((prev) => prev.slice(0, -1));
	};

	const handleCheck = async () => {
		try {
			const result = await PinExists(value);
			nav("/face", { state: { document: result.document } });
		} catch (error) {
			console.error('Erro ao consultar PIN:', error);
		}
	};

		
	return (
    	<div className="flex w-[320px] flex-col justify-start">
    		<div className='px-14 h-[30px] mt-3 flex justify-center items-center'>
				<div className='h-[30px] flex items-center'>
					<input
						type="password"
						className="w-[100px] h-full bg-[#111111] text-gray-100"
						value={value}
						readOnly
						inputMode="numeric"
						minLength={6}
						maxLength={7}
						style={{ letterSpacing: '0.3em' }}
					/>
				</div>
			</div>
    
      		<div className='flex items-center mt-7 ml-10'>
        		<div className="grid grid-cols-3 gap-3 w-[228px]">
          			{[1,2,3,4,5,6,7,8,9].map((num) => (
						<button
							key={num}
							className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]"
							onClick={() => handleClick(num.toString())}
						>
						{num}
						</button>
					))}
					<button
						className={`w-[68px] h-[68px] rounded-md text-white text-2xl font-regular shadow flex items-center justify-center 
						${value.length === 0 ? 'bg-red-200 cursor-not-allowed opacity-50' : 'bg-red-400 hover:bg-[#3a1a1a]'}`}
						onClick={handleDelete}
            			disabled={value.length === 0}
          			>
						<Delete className="size-6" />
          			</button>
          			<button
            			className="w-[68px] h-[68px] rounded-md bg-[#212121] text-white text-2xl font-regular shadow hover:bg-[#292929]"
            			onClick={() => handleClick("0")}
          			>
           	 		0
          			</button>
          			<button
            			className={`w-[68px] h-[68px] rounded-md text-white text-2xl font-regular shadow flex items-center justify-center 
              			${(value.length === 6 || value.length === 7) ? 'bg-green-400 hover:bg-green-680' : 'bg-green-200 cursor-not-allowed opacity-50'}`}
            			onClick={handleCheck}
            			disabled={!(value.length === 6 || value.length === 7)}
          			>
            			<CircleCheckBig className="size-6" />
          			</button>
        		</div>
      		</div>
    	</div>
  	);
}