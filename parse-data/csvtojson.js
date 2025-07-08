import { file, write } from "bun";
import aylar from "./aylar.json";

if (process.argv.length < 3) {
    throw(new Error("Dosya belirtin."));
}

const csvFile = file(process.argv[2]);
let csvData = await csvFile.text();
csvData = csvData.trim().replaceAll("\r", "").split("\n").slice(3);

let times = new Object();

Array.prototype.swapElements = function (x, y) {
    this[x] = this.splice(y, 1, this[x])[0];
};

for (let i = 0; i < csvData.length; i++) {
    /* 
    "yil-ay-gun"
    "2025-03-03" : {
        "İmsak": "05:34",
        "Güneş": "06:59",
        "Öğle": "12:52",
        "İkindi": "16:01",
        "Akşam": "18:35",
        "Yatsı": "19:54"
    }
    */

    const line = csvData[i];
    let dayJson = line.split(",");

    let dayJsonDate = dayJson[0].split(" ").slice(0,-1);
    dayJsonDate[1] = aylar[dayJsonDate[1]];
    dayJsonDate.swapElements(0, 2);

    dayJson[0] = dayJsonDate.join("-");

    times[dayJson[0]] = {
        "İmsak" : dayJson[2],
        "Güneş" : dayJson[3],
        "Öğle"  : dayJson[4],
        "İkindi": dayJson[5],
        "Akşam" : dayJson[6],
        "Yatsı" : dayJson[7]
    }
}

await write(process.argv[2].split(".").slice(0,-1).join(".")+".json", JSON.stringify(times));
