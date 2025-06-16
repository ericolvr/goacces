import { useEffect, useState } from 'react';

export function Clock() {
  const [time, setTime] = useState(new Date());

  useEffect(() => {
    const interval = setInterval(() => {
      setTime(new Date());
    }, 1000);
    return () => clearInterval(interval);
  }, []);

  const hours = String(time.getHours()).padStart(2, '0');
  const minutes = String(time.getMinutes()).padStart(2, '0');
  const seconds = String(time.getSeconds()).padStart(2, '0');

  return (
    <div className="flex items-end justify-center w-full max-w-full flex-wrap gap-x-8 pt-4">
      {/* Horas */}
      <div className="flex flex-col items-center min-w-0 truncate">
        <span className="text-4xl md:text-4xl font-mono font-regular text-white leading-none truncate">{hours}</span>
        <span className="mt-4 text-base md:text-md font-mono tracking-widest text-white/70 uppercase truncate">Hora</span>
      </div>
      {/* Dois pontos */}
      <span className="text-4xl md:text-4xl font-bold text-white mx-2 pb-8">:</span>
      {/* Minutos */}
      <div className="flex flex-col items-center min-w-0 truncate">
        <span className="text-4xl md:text-4xl font-mono font-regular text-white leading-none truncate">{minutes}</span>
        <span className="mt-4 text-base md:text-md font-mono tracking-widest text-white/70 uppercase truncate">Min</span>
      </div>
      {/* Dois pontos */}
      <span className="text-4xl md:text-4xl font-bold text-white mx-2 pb-8">:</span>
      {/* Segundos */}
      <div className="flex flex-col items-center min-w-0 truncate">
        <span className="text-4xl md:text-4xl font-mono font-regular text-white leading-none truncate">{seconds}</span>
        <span className="mt-4 text-base md:text-md font-mono tracking-widest text-white/70 uppercase truncate">Seg</span>
      </div>
    </div>
  );
}
