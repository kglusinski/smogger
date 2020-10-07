import React, {useState, useEffect} from 'react';
import { Bar } from 'react-chartjs-2'
import Select from 'react-select';
import axios from 'axios';

export default function Chart() {
    const [chartData, setChartData] = useState({})
    const [cities, setCities] = useState([])
    const [selectedCity, setSelectedCity] = useState("Warszawa")

    const loadCities = async () => {
        let cities = [];

        const res = await axios.get("http://localhost:8080/v1/cities?country=PL")

        console.log(res);

        for (const city of res.data) {
            cities.push({
                value: city.name,
                label: city.name
            })
        }

        setCities(cities);
    }

    const loadData = async () => {
        let data = [];
        let dates = [];

        console.log(selectedCity)

        const res = await axios.get(`http://localhost:8080/v1/measurements?city=${selectedCity}&date_from=2020-01-01&date_to=2020-03-01`)

        if (res.data.length < 0 ) {
            return
        }

        for (const row of res.data) {
            if (row.parameter === "pm25") {
                data.push(row.value)
                dates.push(row.date.local)
            }

        }

        setChartData({
            labels: dates,
            datasets: [
                {
                    label: `Zanieczyszczenie PM2.5 dla miasta ${selectedCity} [${res.data[0].unit}]`,
                    data: data,
                    borderWidth: 4
                }
            ]
        })

        console.log(res);
    }

    useEffect(() => {
        loadData()
    }, [selectedCity])

    useEffect(() => {
        loadCities()
    }, [])

    return (
        <div>
            <Select
                name="form-field-name"
                options={cities}
                onChange={ e => {setSelectedCity(e.value)}}
            />

            <Bar data={chartData} options={{
                responsive: true,
                title: { text: "Zanieczyszczenie", display: true },
                scales: {
                    yAxes: [
                        {
                            ticks: {
                                autoSkip: true,
                                maxTicksLimit: 10,
                                beginAtZero: true
                            },
                            gridLines: {
                                display: false
                            }
                        }
                    ],
                    xAxes: [
                        {
                            gridLines: {
                                display: false
                            }
                        }
                    ]
                }
            }} />
        </div>
    )
}