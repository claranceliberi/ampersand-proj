export interface Battery {
    ID: number
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: string | null
    Identifier: string
    ModelName: string
    LastSeenOnline: string
    ManufactureDate: string
    BatteryCapacity: number
    ConsumptionRate: number
    SwapsIn: any | null
    SwapsOut: any | null
    BatteryTelematics: any | null
}

export interface Driver {
    ID: number
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: string | null
    FirstName: string
    LastName: string
    PhoneNumber: string
    Email: string
    Swaps: any | null
}

export interface Station {
    ID: number
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: string | null
    Location: string
    Description: string
    Swaps: any | null
}

export interface Agent {
    ID: number
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: string | null
    FirstName: string
    LastName: string
    PhoneNumber: string
    Email: string
    Swaps: any | null
}

export interface Vehicle {
    ID: number
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: string | null
    Identifier: string
    ModelName: string
    ManufactureDate: string
    Description: string
    Swaps: any | null
}

export interface TelemetryData {
    ID: number
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: string | null
    BatteryInID: number
    BatteryInSoc: number
    BatteryOutID: number
    BatteryOutSoc: number
    DriverID: number
    StationID: number
    AgentID: number
    VehicleID: number
    BatteryIn: Battery
    BatteryOut: Battery
    Cost: number
    DistanceCovered: number
    EnergyConsumed: number
    Driver: Driver
    Station: Station
    Agent: Agent
    Vehicle: Vehicle
}

export interface BatteryData {
    ID: number
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: string | null
    BatteryID: number
    Longitude: string
    Latitude: string
    BatterySoc: number
    Charging: boolean
    ChargingRate: number
    Battery: Battery
}
