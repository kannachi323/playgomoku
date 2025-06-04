import {GomokuBoard} from "../components/Board";
import { PlayerBanner } from "../components/Banner";

import { GameProvider } from "../contexts/GameProvider";

export default function Play() {

  async function createGameState() {
    try {
      const response = await fetch('http://localhost:3000/new-game-state', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          p1: {ID: '1', Username: 'Alice'},
          p2: {ID: '2', Username: 'Bob'},
          size: 19,
        })
      })
      const data = await response.json();
      console.log(data);
    } catch (error) {
      //show a popup if there's something wrong
      console.log(error)
    }
  }


  return (
    <GameProvider>
      <div className="w-full h-[90vh] grid grid-cols-17 bg-red-50">
        <div className="col-span-4">
          <PlayerBanner />
        </div>
        <div className="col-span-9 bg-blue-500">
          <GomokuBoard/>
        </div>
        <div className="col-span-4">
          <PlayerBanner />
        </div>
        <div className="col-span-4">
          <button className="border-2 outline-2 w-full"
            onClick={() => createGameState()}
          >
            hello
          </button>
        </div>
        
      </div>
    </GameProvider>
   
  );
}
