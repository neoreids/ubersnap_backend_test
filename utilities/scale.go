package utilities

func ConvertScale(sourceValue int) int {
	// Define the source and target ranges
	sourceMin, sourceMax := 0, 100
	targetMin, targetMax := 0, 9

	// Calculate the ratio between the source range and the target range
	ratio := float64(targetMax - targetMin) / float64(sourceMax - sourceMin)

	// Perform linear interpolation
	targetValue := int(float64(sourceValue - sourceMin) * ratio) + targetMin

	return targetValue
}
