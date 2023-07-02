/**
 * Formats a date string into a readable format.
 * @param dateString a date string of any format
 * @returns {string} A date string in the format of "YYYY-MM-DD HH:MM:SS"
 * @author Clarance Liberi
 */
export function formatTime(dateString: string) {
    const dateObj = new Date(dateString)

    const year = dateObj.getFullYear()
    const month = dateObj.getMonth() + 1 // Month is zero-indexed, so we add 1 to get the correct month number.
    const day = dateObj.getDate()
    const hours = dateObj.getHours()
    const minutes = dateObj.getMinutes()
    const seconds = dateObj.getSeconds()

    const readableTime = `${year}-${month.toString().padStart(2, "0")}-${day.toString().padStart(2, "0")} ${hours.toString().padStart(2, "0")}:${minutes.toString().padStart(2, "0")}:${seconds
        .toString()
        .padStart(2, "0")}`

    return readableTime
}
