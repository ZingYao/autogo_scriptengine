package opencv

import (
	"reflect"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/opencv"
	lua "github.com/yuin/gopher-lua"
)

// OpencvModule opencv 模块
type OpencvModule struct{}

// Name 返回模块名称
func (m *OpencvModule) Name() string {
	return "opencv"
}

// IsAvailable 检查模块是否可用
func (m *OpencvModule) IsAvailable() bool {
	return true
}

type openCVFuncEntry struct {
	name string
	fn   reflect.Value
}

type openCVMethodEntry struct {
	name     string
	goMethod string
}

var openCVFuncs = []openCVFuncEntry{
	{"newPoint2f", reflect.ValueOf(opencv.NewPoint2f)},
	{"newMat", reflect.ValueOf(opencv.NewMat)},
	{"newMatFromCMat", reflect.ValueOf(opencv.NewMatFromCMat)},
	{"newMatWithSize", reflect.ValueOf(opencv.NewMatWithSize)},
	{"newMatWithSizes", reflect.ValueOf(opencv.NewMatWithSizes)},
	{"newMatWithSizesWithScalar", reflect.ValueOf(opencv.NewMatWithSizesWithScalar)},
	{"newMatWithSizesFromBytes", reflect.ValueOf(opencv.NewMatWithSizesFromBytes)},
	{"newMatFromScalar", reflect.ValueOf(opencv.NewMatFromScalar)},
	{"newMatWithSizeFromScalar", reflect.ValueOf(opencv.NewMatWithSizeFromScalar)},
	{"newMatFromBytes", reflect.ValueOf(opencv.NewMatFromBytes)},
	{"eye", reflect.ValueOf(opencv.Eye)},
	{"zeros", reflect.ValueOf(opencv.Zeros)},
	{"ones", reflect.ValueOf(opencv.Ones)},
	{"lUT", reflect.ValueOf(opencv.LUT)},
	{"absDiff", reflect.ValueOf(opencv.AbsDiff)},
	{"add", reflect.ValueOf(opencv.Add)},
	{"addWeighted", reflect.ValueOf(opencv.AddWeighted)},
	{"bitwiseAnd", reflect.ValueOf(opencv.BitwiseAnd)},
	{"bitwiseAndWithMask", reflect.ValueOf(opencv.BitwiseAndWithMask)},
	{"bitwiseNot", reflect.ValueOf(opencv.BitwiseNot)},
	{"bitwiseNotWithMask", reflect.ValueOf(opencv.BitwiseNotWithMask)},
	{"bitwiseOr", reflect.ValueOf(opencv.BitwiseOr)},
	{"bitwiseOrWithMask", reflect.ValueOf(opencv.BitwiseOrWithMask)},
	{"bitwiseXor", reflect.ValueOf(opencv.BitwiseXor)},
	{"bitwiseXorWithMask", reflect.ValueOf(opencv.BitwiseXorWithMask)},
	{"batchDistance", reflect.ValueOf(opencv.BatchDistance)},
	{"borderInterpolate", reflect.ValueOf(opencv.BorderInterpolate)},
	{"calcCovarMatrix", reflect.ValueOf(opencv.CalcCovarMatrix)},
	{"cartToPolar", reflect.ValueOf(opencv.CartToPolar)},
	{"checkRange", reflect.ValueOf(opencv.CheckRange)},
	{"compare", reflect.ValueOf(opencv.Compare)},
	{"countNonZero", reflect.ValueOf(opencv.CountNonZero)},
	{"completeSymm", reflect.ValueOf(opencv.CompleteSymm)},
	{"convertScaleAbs", reflect.ValueOf(opencv.ConvertScaleAbs)},
	{"copyMakeBorder", reflect.ValueOf(opencv.CopyMakeBorder)},
	{"dCT", reflect.ValueOf(opencv.DCT)},
	{"determinant", reflect.ValueOf(opencv.Determinant)},
	{"dFT", reflect.ValueOf(opencv.DFT)},
	{"divide", reflect.ValueOf(opencv.Divide)},
	{"eigen", reflect.ValueOf(opencv.Eigen)},
	{"eigenNonSymmetric", reflect.ValueOf(opencv.EigenNonSymmetric)},
	{"pCABackProject", reflect.ValueOf(opencv.PCABackProject)},
	{"pCACompute", reflect.ValueOf(opencv.PCACompute)},
	{"pCAProject", reflect.ValueOf(opencv.PCAProject)},
	{"pSNR", reflect.ValueOf(opencv.PSNR)},
	{"sVBackSubst", reflect.ValueOf(opencv.SVBackSubst)},
	{"sVDecomp", reflect.ValueOf(opencv.SVDecomp)},
	{"exp", reflect.ValueOf(opencv.Exp)},
	{"extractChannel", reflect.ValueOf(opencv.ExtractChannel)},
	{"findNonZero", reflect.ValueOf(opencv.FindNonZero)},
	{"flip", reflect.ValueOf(opencv.Flip)},
	{"gemm", reflect.ValueOf(opencv.Gemm)},
	{"getOptimalDFTSize", reflect.ValueOf(opencv.GetOptimalDFTSize)},
	{"hconcat", reflect.ValueOf(opencv.Hconcat)},
	{"vconcat", reflect.ValueOf(opencv.Vconcat)},
	{"rotate", reflect.ValueOf(opencv.Rotate)},
	{"iDCT", reflect.ValueOf(opencv.IDCT)},
	{"iDFT", reflect.ValueOf(opencv.IDFT)},
	{"inRange", reflect.ValueOf(opencv.InRange)},
	{"inRangeWithScalar", reflect.ValueOf(opencv.InRangeWithScalar)},
	{"insertChannel", reflect.ValueOf(opencv.InsertChannel)},
	{"invert", reflect.ValueOf(opencv.Invert)},
	{"kMeans", reflect.ValueOf(opencv.KMeans)},
	{"kMeansPoints", reflect.ValueOf(opencv.KMeansPoints)},
	{"log", reflect.ValueOf(opencv.Log)},
	{"magnitude", reflect.ValueOf(opencv.Magnitude)},
	{"mahalanobis", reflect.ValueOf(opencv.Mahalanobis)},
	{"mulTransposed", reflect.ValueOf(opencv.MulTransposed)},
	{"max", reflect.ValueOf(opencv.Max)},
	{"meanStdDev", reflect.ValueOf(opencv.MeanStdDev)},
	{"merge", reflect.ValueOf(opencv.Merge)},
	{"min", reflect.ValueOf(opencv.Min)},
	{"minMaxIdx", reflect.ValueOf(opencv.MinMaxIdx)},
	{"minMaxLoc", reflect.ValueOf(opencv.MinMaxLoc)},
	{"minMaxLocWithMask", reflect.ValueOf(opencv.MinMaxLocWithMask)},
	{"mixChannels", reflect.ValueOf(opencv.MixChannels)},
	{"mulSpectrums", reflect.ValueOf(opencv.MulSpectrums)},
	{"multiply", reflect.ValueOf(opencv.Multiply)},
	{"multiplyWithParams", reflect.ValueOf(opencv.MultiplyWithParams)},
	{"normalize", reflect.ValueOf(opencv.Normalize)},
	{"norm", reflect.ValueOf(opencv.Norm)},
	{"normWithMats", reflect.ValueOf(opencv.NormWithMats)},
	{"perspectiveTransform", reflect.ValueOf(opencv.PerspectiveTransform)},
	{"solve", reflect.ValueOf(opencv.Solve)},
	{"solveCubic", reflect.ValueOf(opencv.SolveCubic)},
	{"solvePoly", reflect.ValueOf(opencv.SolvePoly)},
	{"reduce", reflect.ValueOf(opencv.Reduce)},
	{"reduceArgMax", reflect.ValueOf(opencv.ReduceArgMax)},
	{"reduceArgMin", reflect.ValueOf(opencv.ReduceArgMin)},
	{"repeat", reflect.ValueOf(opencv.Repeat)},
	{"scaleAdd", reflect.ValueOf(opencv.ScaleAdd)},
	{"setIdentity", reflect.ValueOf(opencv.SetIdentity)},
	{"sort", reflect.ValueOf(opencv.Sort)},
	{"sortIdx", reflect.ValueOf(opencv.SortIdx)},
	{"split", reflect.ValueOf(opencv.Split)},
	{"subtract", reflect.ValueOf(opencv.Subtract)},
	{"trace", reflect.ValueOf(opencv.Trace)},
	{"transform", reflect.ValueOf(opencv.Transform)},
	{"transpose", reflect.ValueOf(opencv.Transpose)},
	{"transposeND", reflect.ValueOf(opencv.TransposeND)},
	{"pow", reflect.ValueOf(opencv.Pow)},
	{"polarToCart", reflect.ValueOf(opencv.PolarToCart)},
	{"phase", reflect.ValueOf(opencv.Phase)},
	{"newTermCriteria", reflect.ValueOf(opencv.NewTermCriteria)},
	{"newScalar", reflect.ValueOf(opencv.NewScalar)},
	{"newPointVector", reflect.ValueOf(opencv.NewPointVector)},
	{"newPointVectorFromPoints", reflect.ValueOf(opencv.NewPointVectorFromPoints)},
	{"newPointVectorFromMat", reflect.ValueOf(opencv.NewPointVectorFromMat)},
	{"newPointsVector", reflect.ValueOf(opencv.NewPointsVector)},
	{"newPointsVectorFromPoints", reflect.ValueOf(opencv.NewPointsVectorFromPoints)},
	{"newPoint2fVector", reflect.ValueOf(opencv.NewPoint2fVector)},
	{"newPoint2fVectorFromPoints", reflect.ValueOf(opencv.NewPoint2fVectorFromPoints)},
	{"newPoint2fVectorFromMat", reflect.ValueOf(opencv.NewPoint2fVectorFromMat)},
	{"getTickCount", reflect.ValueOf(opencv.GetTickCount)},
	{"getTickFrequency", reflect.ValueOf(opencv.GetTickFrequency)},
	{"theRNG", reflect.ValueOf(opencv.TheRNG)},
	{"setRNGSeed", reflect.ValueOf(opencv.SetRNGSeed)},
	{"randN", reflect.ValueOf(opencv.RandN)},
	{"randShuffle", reflect.ValueOf(opencv.RandShuffle)},
	{"randShuffleWithParams", reflect.ValueOf(opencv.RandShuffleWithParams)},
	{"randU", reflect.ValueOf(opencv.RandU)},
	{"newPoints2fVector", reflect.ValueOf(opencv.NewPoints2fVector)},
	{"newPoints2fVectorFromPoints", reflect.ValueOf(opencv.NewPoints2fVectorFromPoints)},
	{"newPoint3f", reflect.ValueOf(opencv.NewPoint3f)},
	{"newPoint3fVector", reflect.ValueOf(opencv.NewPoint3fVector)},
	{"newPoint3fVectorFromPoints", reflect.ValueOf(opencv.NewPoint3fVectorFromPoints)},
	{"newPoint3fVectorFromMat", reflect.ValueOf(opencv.NewPoint3fVectorFromMat)},
	{"newPoints3fVector", reflect.ValueOf(opencv.NewPoints3fVector)},
	{"newPoints3fVectorFromPoints", reflect.ValueOf(opencv.NewPoints3fVectorFromPoints)},
	{"setNumThreads", reflect.ValueOf(opencv.SetNumThreads)},
	{"getNumThreads", reflect.ValueOf(opencv.GetNumThreads)},
	{"newRotatedRect", reflect.ValueOf(opencv.NewRotatedRect)},
	{"newRotatedRect2f", reflect.ValueOf(opencv.NewRotatedRect2f)},
	{"lastExceptionError", reflect.ValueOf(opencv.LastExceptionError)},
	{"getLastExceptionMessage", reflect.ValueOf(opencv.GetLastExceptionMessage)},
	{"clearLastException", reflect.ValueOf(opencv.ClearLastException)},
	{"getLastException", reflect.ValueOf(opencv.GetLastException)},
	{"openCVResult", reflect.ValueOf(opencv.OpenCVResult)},
	{"iMRead", reflect.ValueOf(opencv.IMRead)},
	{"iMReadMulti", reflect.ValueOf(opencv.IMReadMulti)},
	{"iMReadMulti_WithParams", reflect.ValueOf(opencv.IMReadMulti_WithParams)},
	{"iMWrite", reflect.ValueOf(opencv.IMWrite)},
	{"iMWriteWithParams", reflect.ValueOf(opencv.IMWriteWithParams)},
	{"iMEncode", reflect.ValueOf(opencv.IMEncode)},
	{"iMEncodeWithParams", reflect.ValueOf(opencv.IMEncodeWithParams)},
	{"iMDecode", reflect.ValueOf(opencv.IMDecode)},
	{"iMDecodeIntoMat", reflect.ValueOf(opencv.IMDecodeIntoMat)},
	{"arcLength", reflect.ValueOf(opencv.ArcLength)},
	{"approxPolyDP", reflect.ValueOf(opencv.ApproxPolyDP)},
	{"convexHull", reflect.ValueOf(opencv.ConvexHull)},
	{"convexityDefects", reflect.ValueOf(opencv.ConvexityDefects)},
	{"cvtColor", reflect.ValueOf(opencv.CvtColor)},
	{"demosaicing", reflect.ValueOf(opencv.Demosaicing)},
	{"equalizeHist", reflect.ValueOf(opencv.EqualizeHist)},
	{"calcHist", reflect.ValueOf(opencv.CalcHist)},
	{"calcBackProject", reflect.ValueOf(opencv.CalcBackProject)},
	{"compareHist", reflect.ValueOf(opencv.CompareHist)},
	{"eMD", reflect.ValueOf(opencv.EMD)},
	{"clipLine", reflect.ValueOf(opencv.ClipLine)},
	{"bilateralFilter", reflect.ValueOf(opencv.BilateralFilter)},
	{"blur", reflect.ValueOf(opencv.Blur)},
	{"boxFilter", reflect.ValueOf(opencv.BoxFilter)},
	{"sqBoxFilter", reflect.ValueOf(opencv.SqBoxFilter)},
	{"dilate", reflect.ValueOf(opencv.Dilate)},
	{"dilateWithParams", reflect.ValueOf(opencv.DilateWithParams)},
	{"distanceTransform", reflect.ValueOf(opencv.DistanceTransform)},
	{"erode", reflect.ValueOf(opencv.Erode)},
	{"erodeWithParams", reflect.ValueOf(opencv.ErodeWithParams)},
	{"erodeWithParamsAndBorderValue", reflect.ValueOf(opencv.ErodeWithParamsAndBorderValue)},
	{"boundingRect", reflect.ValueOf(opencv.BoundingRect)},
	{"boxPoints", reflect.ValueOf(opencv.BoxPoints)},
	{"boxPoints2f", reflect.ValueOf(opencv.BoxPoints2f)},
	{"contourArea", reflect.ValueOf(opencv.ContourArea)},
	{"minAreaRect", reflect.ValueOf(opencv.MinAreaRect)},
	{"minAreaRect2f", reflect.ValueOf(opencv.MinAreaRect2f)},
	{"fitEllipse", reflect.ValueOf(opencv.FitEllipse)},
	{"minEnclosingCircle", reflect.ValueOf(opencv.MinEnclosingCircle)},
	{"findContours", reflect.ValueOf(opencv.FindContours)},
	{"findContoursWithParams", reflect.ValueOf(opencv.FindContoursWithParams)},
	{"pointPolygonTest", reflect.ValueOf(opencv.PointPolygonTest)},
	{"connectedComponents", reflect.ValueOf(opencv.ConnectedComponents)},
	{"connectedComponentsWithParams", reflect.ValueOf(opencv.ConnectedComponentsWithParams)},
	{"connectedComponentsWithStats", reflect.ValueOf(opencv.ConnectedComponentsWithStats)},
	{"connectedComponentsWithStatsWithParams", reflect.ValueOf(opencv.ConnectedComponentsWithStatsWithParams)},
	{"matchTemplate", reflect.ValueOf(opencv.MatchTemplate)},
	{"moments", reflect.ValueOf(opencv.Moments)},
	{"pyrDown", reflect.ValueOf(opencv.PyrDown)},
	{"pyrUp", reflect.ValueOf(opencv.PyrUp)},
	{"morphologyDefaultBorderValue", reflect.ValueOf(opencv.MorphologyDefaultBorderValue)},
	{"morphologyEx", reflect.ValueOf(opencv.MorphologyEx)},
	{"morphologyExWithParams", reflect.ValueOf(opencv.MorphologyExWithParams)},
	{"getStructuringElement", reflect.ValueOf(opencv.GetStructuringElement)},
	{"gaussianBlur", reflect.ValueOf(opencv.GaussianBlur)},
	{"getGaussianKernel", reflect.ValueOf(opencv.GetGaussianKernel)},
	{"getGaussianKernelWithParams", reflect.ValueOf(opencv.GetGaussianKernelWithParams)},
	{"sobel", reflect.ValueOf(opencv.Sobel)},
	{"spatialGradient", reflect.ValueOf(opencv.SpatialGradient)},
	{"laplacian", reflect.ValueOf(opencv.Laplacian)},
	{"scharr", reflect.ValueOf(opencv.Scharr)},
	{"medianBlur", reflect.ValueOf(opencv.MedianBlur)},
	{"canny", reflect.ValueOf(opencv.Canny)},
	{"cornerSubPix", reflect.ValueOf(opencv.CornerSubPix)},
	{"goodFeaturesToTrack", reflect.ValueOf(opencv.GoodFeaturesToTrack)},
	{"grabCut", reflect.ValueOf(opencv.GrabCut)},
	{"houghCircles", reflect.ValueOf(opencv.HoughCircles)},
	{"houghCirclesWithParams", reflect.ValueOf(opencv.HoughCirclesWithParams)},
	{"houghLines", reflect.ValueOf(opencv.HoughLines)},
	{"houghLinesP", reflect.ValueOf(opencv.HoughLinesP)},
	{"houghLinesPWithParams", reflect.ValueOf(opencv.HoughLinesPWithParams)},
	{"houghLinesPointSet", reflect.ValueOf(opencv.HoughLinesPointSet)},
	{"integral", reflect.ValueOf(opencv.Integral)},
	{"threshold", reflect.ValueOf(opencv.Threshold)},
	{"adaptiveThreshold", reflect.ValueOf(opencv.AdaptiveThreshold)},
	{"arrowedLine", reflect.ValueOf(opencv.ArrowedLine)},
	{"circle", reflect.ValueOf(opencv.Circle)},
	{"circleWithParams", reflect.ValueOf(opencv.CircleWithParams)},
	{"ellipse", reflect.ValueOf(opencv.Ellipse)},
	{"ellipseWithParams", reflect.ValueOf(opencv.EllipseWithParams)},
	{"line", reflect.ValueOf(opencv.Line)},
	{"rectangle", reflect.ValueOf(opencv.Rectangle)},
	{"rectangleWithParams", reflect.ValueOf(opencv.RectangleWithParams)},
	{"fillPoly", reflect.ValueOf(opencv.FillPoly)},
	{"fillPolyWithParams", reflect.ValueOf(opencv.FillPolyWithParams)},
	{"polylines", reflect.ValueOf(opencv.Polylines)},
	{"getTextSize", reflect.ValueOf(opencv.GetTextSize)},
	{"getTextSizeWithBaseline", reflect.ValueOf(opencv.GetTextSizeWithBaseline)},
	{"putText", reflect.ValueOf(opencv.PutText)},
	{"putTextWithParams", reflect.ValueOf(opencv.PutTextWithParams)},
	{"resize", reflect.ValueOf(opencv.Resize)},
	{"getRectSubPix", reflect.ValueOf(opencv.GetRectSubPix)},
	{"getRotationMatrix2D", reflect.ValueOf(opencv.GetRotationMatrix2D)},
	{"warpAffine", reflect.ValueOf(opencv.WarpAffine)},
	{"warpAffineWithParams", reflect.ValueOf(opencv.WarpAffineWithParams)},
	{"warpPerspective", reflect.ValueOf(opencv.WarpPerspective)},
	{"warpPerspectiveWithParams", reflect.ValueOf(opencv.WarpPerspectiveWithParams)},
	{"watershed", reflect.ValueOf(opencv.Watershed)},
	{"applyColorMap", reflect.ValueOf(opencv.ApplyColorMap)},
	{"applyCustomColorMap", reflect.ValueOf(opencv.ApplyCustomColorMap)},
	{"getPerspectiveTransform", reflect.ValueOf(opencv.GetPerspectiveTransform)},
	{"getPerspectiveTransform2f", reflect.ValueOf(opencv.GetPerspectiveTransform2f)},
	{"getAffineTransform", reflect.ValueOf(opencv.GetAffineTransform)},
	{"getAffineTransform2f", reflect.ValueOf(opencv.GetAffineTransform2f)},
	{"drawContours", reflect.ValueOf(opencv.DrawContours)},
	{"drawContoursWithParams", reflect.ValueOf(opencv.DrawContoursWithParams)},
	{"remap", reflect.ValueOf(opencv.Remap)},
	{"filter2D", reflect.ValueOf(opencv.Filter2D)},
	{"sepFilter2D", reflect.ValueOf(opencv.SepFilter2D)},
	{"logPolar", reflect.ValueOf(opencv.LogPolar)},
	{"linearPolar", reflect.ValueOf(opencv.LinearPolar)},
	{"fitLine", reflect.ValueOf(opencv.FitLine)},
	{"matchShapes", reflect.ValueOf(opencv.MatchShapes)},
	{"newCLAHE", reflect.ValueOf(opencv.NewCLAHE)},
	{"newCLAHEWithParams", reflect.ValueOf(opencv.NewCLAHEWithParams)},
	{"invertAffineTransform", reflect.ValueOf(opencv.InvertAffineTransform)},
	{"phaseCorrelate", reflect.ValueOf(opencv.PhaseCorrelate)},
	{"createHanningWindow", reflect.ValueOf(opencv.CreateHanningWindow)},
	{"imageToMatRGBA", reflect.ValueOf(opencv.ImageToMatRGBA)},
	{"imageToMatRGB", reflect.ValueOf(opencv.ImageToMatRGB)},
	{"imageGrayToMatGray", reflect.ValueOf(opencv.ImageGrayToMatGray)},
	{"accumulate", reflect.ValueOf(opencv.Accumulate)},
	{"accumulateWithMask", reflect.ValueOf(opencv.AccumulateWithMask)},
	{"accumulateSquare", reflect.ValueOf(opencv.AccumulateSquare)},
	{"accumulateSquareWithMask", reflect.ValueOf(opencv.AccumulateSquareWithMask)},
	{"accumulateProduct", reflect.ValueOf(opencv.AccumulateProduct)},
	{"accumulateProductWithMask", reflect.ValueOf(opencv.AccumulateProductWithMask)},
	{"accumulatedWeighted", reflect.ValueOf(opencv.AccumulatedWeighted)},
	{"accumulatedWeightedWithMask", reflect.ValueOf(opencv.AccumulatedWeightedWithMask)},
	{"findImage", reflect.ValueOf(opencv.FindImage)},
	{"findImageAll", reflect.ValueOf(opencv.FindImageAll)},
}

var openCVMethods = map[string][]openCVMethodEntry{
	"Mat": {
		{"close", "Close"},
		{"fromPtr", "FromPtr"},
		{"ptr", "Ptr"},
		{"empty", "Empty"},
		{"closed", "Closed"},
		{"isContinuous", "IsContinuous"},
		{"inv", "Inv"},
		{"col", "Col"},
		{"row", "Row"},
		{"clone", "Clone"},
		{"copyTo", "CopyTo"},
		{"copyToWithMask", "CopyToWithMask"},
		{"convertTo", "ConvertTo"},
		{"convertToWithParams", "ConvertToWithParams"},
		{"total", "Total"},
		{"size", "Size"},
		{"toBytes", "ToBytes"},
		{"dataPtrUint8", "DataPtrUint8"},
		{"dataPtrInt8", "DataPtrInt8"},
		{"dataPtrUint16", "DataPtrUint16"},
		{"dataPtrInt16", "DataPtrInt16"},
		{"dataPtrFloat32", "DataPtrFloat32"},
		{"dataPtrFloat64", "DataPtrFloat64"},
		{"region", "Region"},
		{"reshape", "Reshape"},
		{"convertFp16", "ConvertFp16"},
		{"mean", "Mean"},
		{"meanWithMask", "MeanWithMask"},
		{"sqrt", "Sqrt"},
		{"sum", "Sum"},
		{"patchNaNs", "PatchNaNs"},
		{"rows", "Rows"},
		{"cols", "Cols"},
		{"channels", "Channels"},
		{"type", "Type"},
		{"step", "Step"},
		{"elemSize", "ElemSize"},
		{"getUCharAt", "GetUCharAt"},
		{"getUCharAt3", "GetUCharAt3"},
		{"getSCharAt", "GetSCharAt"},
		{"getSCharAt3", "GetSCharAt3"},
		{"getShortAt", "GetShortAt"},
		{"getShortAt3", "GetShortAt3"},
		{"getIntAt", "GetIntAt"},
		{"getIntAt3", "GetIntAt3"},
		{"getFloatAt", "GetFloatAt"},
		{"getFloatAt3", "GetFloatAt3"},
		{"getDoubleAt", "GetDoubleAt"},
		{"getDoubleAt3", "GetDoubleAt3"},
		{"setTo", "SetTo"},
		{"setUCharAt", "SetUCharAt"},
		{"setUCharAt3", "SetUCharAt3"},
		{"setSCharAt", "SetSCharAt"},
		{"setSCharAt3", "SetSCharAt3"},
		{"setShortAt", "SetShortAt"},
		{"setShortAt3", "SetShortAt3"},
		{"setIntAt", "SetIntAt"},
		{"setIntAt3", "SetIntAt3"},
		{"setFloatAt", "SetFloatAt"},
		{"setFloatAt3", "SetFloatAt3"},
		{"setDoubleAt", "SetDoubleAt"},
		{"setDoubleAt3", "SetDoubleAt3"},
		{"addUChar", "AddUChar"},
		{"subtractUChar", "SubtractUChar"},
		{"multiplyUChar", "MultiplyUChar"},
		{"divideUChar", "DivideUChar"},
		{"addFloat", "AddFloat"},
		{"subtractFloat", "SubtractFloat"},
		{"multiplyFloat", "MultiplyFloat"},
		{"divideFloat", "DivideFloat"},
		{"multiplyMatrix", "MultiplyMatrix"},
		{"t", "T"},
		{"getVecbAt", "GetVecbAt"},
		{"getVecfAt", "GetVecfAt"},
		{"getVecdAt", "GetVecdAt"},
		{"getVeciAt", "GetVeciAt"},
		{"rowRange", "RowRange"},
		{"colRange", "ColRange"},
		{"toImage", "ToImage"},
		{"toImageYUV", "ToImageYUV"},
		{"toImageYUVWithParams", "ToImageYUVWithParams"},
	},
	"PointVector": {
		{"isNil", "IsNil"},
		{"size", "Size"},
		{"at", "At"},
		{"append", "Append"},
		{"toPoints", "ToPoints"},
		{"close", "Close"},
	},
	"PointsVector": {
		{"p", "P"},
		{"toPoints", "ToPoints"},
		{"isNil", "IsNil"},
		{"size", "Size"},
		{"at", "At"},
		{"append", "Append"},
		{"close", "Close"},
	},
	"Point2fVector": {
		{"isNil", "IsNil"},
		{"size", "Size"},
		{"at", "At"},
		{"toPoints", "ToPoints"},
		{"close", "Close"},
	},
	"RNG": {
		{"fill", "Fill"},
		{"gaussian", "Gaussian"},
		{"next", "Next"},
	},
	"NativeByteBuffer": {
		{"getBytes", "GetBytes"},
		{"len", "Len"},
		{"close", "Close"},
	},
	"Points2fVector": {
		{"p", "P"},
		{"toPoints", "ToPoints"},
		{"isNil", "IsNil"},
		{"size", "Size"},
		{"at", "At"},
		{"append", "Append"},
		{"close", "Close"},
	},
	"Point3fVector": {
		{"isNil", "IsNil"},
		{"size", "Size"},
		{"at", "At"},
		{"append", "Append"},
		{"toPoints", "ToPoints"},
		{"close", "Close"},
	},
	"Points3fVector": {
		{"toPoints", "ToPoints"},
		{"isNil", "IsNil"},
		{"size", "Size"},
		{"at", "At"},
		{"append", "Append"},
		{"close", "Close"},
	},
	"CLAHE": {
		{"close", "Close"},
		{"apply", "Apply"},
	},
}

// Register 向引擎注册方法
func (m *OpencvModule) Register(engine model.Engine) error {
	state := engine.GetState()

	opencvObj := state.NewTable()
	state.SetGlobal("opencv", opencvObj)
	for _, entry := range openCVFuncs {
		opencvObj.RawSetString(entry.name, state.NewFunction(bindOpenCVLuaFunc(entry.fn)))
	}

	engine.RegisterMethod("opencv.newPoint2f", "OpenCV.NewPoint2f", opencv.NewPoint2f, true)
	engine.RegisterMethod("opencv.newMat", "OpenCV.NewMat", opencv.NewMat, true)
	engine.RegisterMethod("opencv.newMatFromCMat", "OpenCV.NewMatFromCMat", opencv.NewMatFromCMat, true)
	engine.RegisterMethod("opencv.newMatWithSize", "OpenCV.NewMatWithSize", opencv.NewMatWithSize, true)
	engine.RegisterMethod("opencv.newMatWithSizes", "OpenCV.NewMatWithSizes", opencv.NewMatWithSizes, true)
	engine.RegisterMethod("opencv.newMatWithSizesWithScalar", "OpenCV.NewMatWithSizesWithScalar", opencv.NewMatWithSizesWithScalar, true)
	engine.RegisterMethod("opencv.newMatWithSizesFromBytes", "OpenCV.NewMatWithSizesFromBytes", opencv.NewMatWithSizesFromBytes, true)
	engine.RegisterMethod("opencv.newMatFromScalar", "OpenCV.NewMatFromScalar", opencv.NewMatFromScalar, true)
	engine.RegisterMethod("opencv.newMatWithSizeFromScalar", "OpenCV.NewMatWithSizeFromScalar", opencv.NewMatWithSizeFromScalar, true)
	engine.RegisterMethod("opencv.newMatFromBytes", "OpenCV.NewMatFromBytes", opencv.NewMatFromBytes, true)
	engine.RegisterMethod("opencv.eye", "OpenCV.Eye", opencv.Eye, true)
	engine.RegisterMethod("opencv.zeros", "OpenCV.Zeros", opencv.Zeros, true)
	engine.RegisterMethod("opencv.ones", "OpenCV.Ones", opencv.Ones, true)
	engine.RegisterMethod("opencv.lUT", "OpenCV.LUT", opencv.LUT, true)
	engine.RegisterMethod("opencv.absDiff", "OpenCV.AbsDiff", opencv.AbsDiff, true)
	engine.RegisterMethod("opencv.add", "OpenCV.Add", opencv.Add, true)
	engine.RegisterMethod("opencv.addWeighted", "OpenCV.AddWeighted", opencv.AddWeighted, true)
	engine.RegisterMethod("opencv.bitwiseAnd", "OpenCV.BitwiseAnd", opencv.BitwiseAnd, true)
	engine.RegisterMethod("opencv.bitwiseAndWithMask", "OpenCV.BitwiseAndWithMask", opencv.BitwiseAndWithMask, true)
	engine.RegisterMethod("opencv.bitwiseNot", "OpenCV.BitwiseNot", opencv.BitwiseNot, true)
	engine.RegisterMethod("opencv.bitwiseNotWithMask", "OpenCV.BitwiseNotWithMask", opencv.BitwiseNotWithMask, true)
	engine.RegisterMethod("opencv.bitwiseOr", "OpenCV.BitwiseOr", opencv.BitwiseOr, true)
	engine.RegisterMethod("opencv.bitwiseOrWithMask", "OpenCV.BitwiseOrWithMask", opencv.BitwiseOrWithMask, true)
	engine.RegisterMethod("opencv.bitwiseXor", "OpenCV.BitwiseXor", opencv.BitwiseXor, true)
	engine.RegisterMethod("opencv.bitwiseXorWithMask", "OpenCV.BitwiseXorWithMask", opencv.BitwiseXorWithMask, true)
	engine.RegisterMethod("opencv.batchDistance", "OpenCV.BatchDistance", opencv.BatchDistance, true)
	engine.RegisterMethod("opencv.borderInterpolate", "OpenCV.BorderInterpolate", opencv.BorderInterpolate, true)
	engine.RegisterMethod("opencv.calcCovarMatrix", "OpenCV.CalcCovarMatrix", opencv.CalcCovarMatrix, true)
	engine.RegisterMethod("opencv.cartToPolar", "OpenCV.CartToPolar", opencv.CartToPolar, true)
	engine.RegisterMethod("opencv.checkRange", "OpenCV.CheckRange", opencv.CheckRange, true)
	engine.RegisterMethod("opencv.compare", "OpenCV.Compare", opencv.Compare, true)
	engine.RegisterMethod("opencv.countNonZero", "OpenCV.CountNonZero", opencv.CountNonZero, true)
	engine.RegisterMethod("opencv.completeSymm", "OpenCV.CompleteSymm", opencv.CompleteSymm, true)
	engine.RegisterMethod("opencv.convertScaleAbs", "OpenCV.ConvertScaleAbs", opencv.ConvertScaleAbs, true)
	engine.RegisterMethod("opencv.copyMakeBorder", "OpenCV.CopyMakeBorder", opencv.CopyMakeBorder, true)
	engine.RegisterMethod("opencv.dCT", "OpenCV.DCT", opencv.DCT, true)
	engine.RegisterMethod("opencv.determinant", "OpenCV.Determinant", opencv.Determinant, true)
	engine.RegisterMethod("opencv.dFT", "OpenCV.DFT", opencv.DFT, true)
	engine.RegisterMethod("opencv.divide", "OpenCV.Divide", opencv.Divide, true)
	engine.RegisterMethod("opencv.eigen", "OpenCV.Eigen", opencv.Eigen, true)
	engine.RegisterMethod("opencv.eigenNonSymmetric", "OpenCV.EigenNonSymmetric", opencv.EigenNonSymmetric, true)
	engine.RegisterMethod("opencv.pCABackProject", "OpenCV.PCABackProject", opencv.PCABackProject, true)
	engine.RegisterMethod("opencv.pCACompute", "OpenCV.PCACompute", opencv.PCACompute, true)
	engine.RegisterMethod("opencv.pCAProject", "OpenCV.PCAProject", opencv.PCAProject, true)
	engine.RegisterMethod("opencv.pSNR", "OpenCV.PSNR", opencv.PSNR, true)
	engine.RegisterMethod("opencv.sVBackSubst", "OpenCV.SVBackSubst", opencv.SVBackSubst, true)
	engine.RegisterMethod("opencv.sVDecomp", "OpenCV.SVDecomp", opencv.SVDecomp, true)
	engine.RegisterMethod("opencv.exp", "OpenCV.Exp", opencv.Exp, true)
	engine.RegisterMethod("opencv.extractChannel", "OpenCV.ExtractChannel", opencv.ExtractChannel, true)
	engine.RegisterMethod("opencv.findNonZero", "OpenCV.FindNonZero", opencv.FindNonZero, true)
	engine.RegisterMethod("opencv.flip", "OpenCV.Flip", opencv.Flip, true)
	engine.RegisterMethod("opencv.gemm", "OpenCV.Gemm", opencv.Gemm, true)
	engine.RegisterMethod("opencv.getOptimalDFTSize", "OpenCV.GetOptimalDFTSize", opencv.GetOptimalDFTSize, true)
	engine.RegisterMethod("opencv.hconcat", "OpenCV.Hconcat", opencv.Hconcat, true)
	engine.RegisterMethod("opencv.vconcat", "OpenCV.Vconcat", opencv.Vconcat, true)
	engine.RegisterMethod("opencv.rotate", "OpenCV.Rotate", opencv.Rotate, true)
	engine.RegisterMethod("opencv.iDCT", "OpenCV.IDCT", opencv.IDCT, true)
	engine.RegisterMethod("opencv.iDFT", "OpenCV.IDFT", opencv.IDFT, true)
	engine.RegisterMethod("opencv.inRange", "OpenCV.InRange", opencv.InRange, true)
	engine.RegisterMethod("opencv.inRangeWithScalar", "OpenCV.InRangeWithScalar", opencv.InRangeWithScalar, true)
	engine.RegisterMethod("opencv.insertChannel", "OpenCV.InsertChannel", opencv.InsertChannel, true)
	engine.RegisterMethod("opencv.invert", "OpenCV.Invert", opencv.Invert, true)
	engine.RegisterMethod("opencv.kMeans", "OpenCV.KMeans", opencv.KMeans, true)
	engine.RegisterMethod("opencv.kMeansPoints", "OpenCV.KMeansPoints", opencv.KMeansPoints, true)
	engine.RegisterMethod("opencv.log", "OpenCV.Log", opencv.Log, true)
	engine.RegisterMethod("opencv.magnitude", "OpenCV.Magnitude", opencv.Magnitude, true)
	engine.RegisterMethod("opencv.mahalanobis", "OpenCV.Mahalanobis", opencv.Mahalanobis, true)
	engine.RegisterMethod("opencv.mulTransposed", "OpenCV.MulTransposed", opencv.MulTransposed, true)
	engine.RegisterMethod("opencv.max", "OpenCV.Max", opencv.Max, true)
	engine.RegisterMethod("opencv.meanStdDev", "OpenCV.MeanStdDev", opencv.MeanStdDev, true)
	engine.RegisterMethod("opencv.merge", "OpenCV.Merge", opencv.Merge, true)
	engine.RegisterMethod("opencv.min", "OpenCV.Min", opencv.Min, true)
	engine.RegisterMethod("opencv.minMaxIdx", "OpenCV.MinMaxIdx", opencv.MinMaxIdx, true)
	engine.RegisterMethod("opencv.minMaxLoc", "OpenCV.MinMaxLoc", opencv.MinMaxLoc, true)
	engine.RegisterMethod("opencv.minMaxLocWithMask", "OpenCV.MinMaxLocWithMask", opencv.MinMaxLocWithMask, true)
	engine.RegisterMethod("opencv.mixChannels", "OpenCV.MixChannels", opencv.MixChannels, true)
	engine.RegisterMethod("opencv.mulSpectrums", "OpenCV.MulSpectrums", opencv.MulSpectrums, true)
	engine.RegisterMethod("opencv.multiply", "OpenCV.Multiply", opencv.Multiply, true)
	engine.RegisterMethod("opencv.multiplyWithParams", "OpenCV.MultiplyWithParams", opencv.MultiplyWithParams, true)
	engine.RegisterMethod("opencv.normalize", "OpenCV.Normalize", opencv.Normalize, true)
	engine.RegisterMethod("opencv.norm", "OpenCV.Norm", opencv.Norm, true)
	engine.RegisterMethod("opencv.normWithMats", "OpenCV.NormWithMats", opencv.NormWithMats, true)
	engine.RegisterMethod("opencv.perspectiveTransform", "OpenCV.PerspectiveTransform", opencv.PerspectiveTransform, true)
	engine.RegisterMethod("opencv.solve", "OpenCV.Solve", opencv.Solve, true)
	engine.RegisterMethod("opencv.solveCubic", "OpenCV.SolveCubic", opencv.SolveCubic, true)
	engine.RegisterMethod("opencv.solvePoly", "OpenCV.SolvePoly", opencv.SolvePoly, true)
	engine.RegisterMethod("opencv.reduce", "OpenCV.Reduce", opencv.Reduce, true)
	engine.RegisterMethod("opencv.reduceArgMax", "OpenCV.ReduceArgMax", opencv.ReduceArgMax, true)
	engine.RegisterMethod("opencv.reduceArgMin", "OpenCV.ReduceArgMin", opencv.ReduceArgMin, true)
	engine.RegisterMethod("opencv.repeat", "OpenCV.Repeat", opencv.Repeat, true)
	engine.RegisterMethod("opencv.scaleAdd", "OpenCV.ScaleAdd", opencv.ScaleAdd, true)
	engine.RegisterMethod("opencv.setIdentity", "OpenCV.SetIdentity", opencv.SetIdentity, true)
	engine.RegisterMethod("opencv.sort", "OpenCV.Sort", opencv.Sort, true)
	engine.RegisterMethod("opencv.sortIdx", "OpenCV.SortIdx", opencv.SortIdx, true)
	engine.RegisterMethod("opencv.split", "OpenCV.Split", opencv.Split, true)
	engine.RegisterMethod("opencv.subtract", "OpenCV.Subtract", opencv.Subtract, true)
	engine.RegisterMethod("opencv.trace", "OpenCV.Trace", opencv.Trace, true)
	engine.RegisterMethod("opencv.transform", "OpenCV.Transform", opencv.Transform, true)
	engine.RegisterMethod("opencv.transpose", "OpenCV.Transpose", opencv.Transpose, true)
	engine.RegisterMethod("opencv.transposeND", "OpenCV.TransposeND", opencv.TransposeND, true)
	engine.RegisterMethod("opencv.pow", "OpenCV.Pow", opencv.Pow, true)
	engine.RegisterMethod("opencv.polarToCart", "OpenCV.PolarToCart", opencv.PolarToCart, true)
	engine.RegisterMethod("opencv.phase", "OpenCV.Phase", opencv.Phase, true)
	engine.RegisterMethod("opencv.newTermCriteria", "OpenCV.NewTermCriteria", opencv.NewTermCriteria, true)
	engine.RegisterMethod("opencv.newScalar", "OpenCV.NewScalar", opencv.NewScalar, true)
	engine.RegisterMethod("opencv.newPointVector", "OpenCV.NewPointVector", opencv.NewPointVector, true)
	engine.RegisterMethod("opencv.newPointVectorFromPoints", "OpenCV.NewPointVectorFromPoints", opencv.NewPointVectorFromPoints, true)
	engine.RegisterMethod("opencv.newPointVectorFromMat", "OpenCV.NewPointVectorFromMat", opencv.NewPointVectorFromMat, true)
	engine.RegisterMethod("opencv.newPointsVector", "OpenCV.NewPointsVector", opencv.NewPointsVector, true)
	engine.RegisterMethod("opencv.newPointsVectorFromPoints", "OpenCV.NewPointsVectorFromPoints", opencv.NewPointsVectorFromPoints, true)
	engine.RegisterMethod("opencv.newPoint2fVector", "OpenCV.NewPoint2fVector", opencv.NewPoint2fVector, true)
	engine.RegisterMethod("opencv.newPoint2fVectorFromPoints", "OpenCV.NewPoint2fVectorFromPoints", opencv.NewPoint2fVectorFromPoints, true)
	engine.RegisterMethod("opencv.newPoint2fVectorFromMat", "OpenCV.NewPoint2fVectorFromMat", opencv.NewPoint2fVectorFromMat, true)
	engine.RegisterMethod("opencv.getTickCount", "OpenCV.GetTickCount", opencv.GetTickCount, true)
	engine.RegisterMethod("opencv.getTickFrequency", "OpenCV.GetTickFrequency", opencv.GetTickFrequency, true)
	engine.RegisterMethod("opencv.theRNG", "OpenCV.TheRNG", opencv.TheRNG, true)
	engine.RegisterMethod("opencv.setRNGSeed", "OpenCV.SetRNGSeed", opencv.SetRNGSeed, true)
	engine.RegisterMethod("opencv.randN", "OpenCV.RandN", opencv.RandN, true)
	engine.RegisterMethod("opencv.randShuffle", "OpenCV.RandShuffle", opencv.RandShuffle, true)
	engine.RegisterMethod("opencv.randShuffleWithParams", "OpenCV.RandShuffleWithParams", opencv.RandShuffleWithParams, true)
	engine.RegisterMethod("opencv.randU", "OpenCV.RandU", opencv.RandU, true)
	engine.RegisterMethod("opencv.newPoints2fVector", "OpenCV.NewPoints2fVector", opencv.NewPoints2fVector, true)
	engine.RegisterMethod("opencv.newPoints2fVectorFromPoints", "OpenCV.NewPoints2fVectorFromPoints", opencv.NewPoints2fVectorFromPoints, true)
	engine.RegisterMethod("opencv.newPoint3f", "OpenCV.NewPoint3f", opencv.NewPoint3f, true)
	engine.RegisterMethod("opencv.newPoint3fVector", "OpenCV.NewPoint3fVector", opencv.NewPoint3fVector, true)
	engine.RegisterMethod("opencv.newPoint3fVectorFromPoints", "OpenCV.NewPoint3fVectorFromPoints", opencv.NewPoint3fVectorFromPoints, true)
	engine.RegisterMethod("opencv.newPoint3fVectorFromMat", "OpenCV.NewPoint3fVectorFromMat", opencv.NewPoint3fVectorFromMat, true)
	engine.RegisterMethod("opencv.newPoints3fVector", "OpenCV.NewPoints3fVector", opencv.NewPoints3fVector, true)
	engine.RegisterMethod("opencv.newPoints3fVectorFromPoints", "OpenCV.NewPoints3fVectorFromPoints", opencv.NewPoints3fVectorFromPoints, true)
	engine.RegisterMethod("opencv.setNumThreads", "OpenCV.SetNumThreads", opencv.SetNumThreads, true)
	engine.RegisterMethod("opencv.getNumThreads", "OpenCV.GetNumThreads", opencv.GetNumThreads, true)
	engine.RegisterMethod("opencv.newRotatedRect", "OpenCV.NewRotatedRect", opencv.NewRotatedRect, true)
	engine.RegisterMethod("opencv.newRotatedRect2f", "OpenCV.NewRotatedRect2f", opencv.NewRotatedRect2f, true)
	engine.RegisterMethod("opencv.lastExceptionError", "OpenCV.LastExceptionError", opencv.LastExceptionError, true)
	engine.RegisterMethod("opencv.getLastExceptionMessage", "OpenCV.GetLastExceptionMessage", opencv.GetLastExceptionMessage, true)
	engine.RegisterMethod("opencv.clearLastException", "OpenCV.ClearLastException", opencv.ClearLastException, true)
	engine.RegisterMethod("opencv.getLastException", "OpenCV.GetLastException", opencv.GetLastException, true)
	engine.RegisterMethod("opencv.openCVResult", "OpenCV.OpenCVResult", opencv.OpenCVResult, true)
	engine.RegisterMethod("opencv.iMRead", "OpenCV.IMRead", opencv.IMRead, true)
	engine.RegisterMethod("opencv.iMReadMulti", "OpenCV.IMReadMulti", opencv.IMReadMulti, true)
	engine.RegisterMethod("opencv.iMReadMulti_WithParams", "OpenCV.IMReadMulti_WithParams", opencv.IMReadMulti_WithParams, true)
	engine.RegisterMethod("opencv.iMWrite", "OpenCV.IMWrite", opencv.IMWrite, true)
	engine.RegisterMethod("opencv.iMWriteWithParams", "OpenCV.IMWriteWithParams", opencv.IMWriteWithParams, true)
	engine.RegisterMethod("opencv.iMEncode", "OpenCV.IMEncode", opencv.IMEncode, true)
	engine.RegisterMethod("opencv.iMEncodeWithParams", "OpenCV.IMEncodeWithParams", opencv.IMEncodeWithParams, true)
	engine.RegisterMethod("opencv.iMDecode", "OpenCV.IMDecode", opencv.IMDecode, true)
	engine.RegisterMethod("opencv.iMDecodeIntoMat", "OpenCV.IMDecodeIntoMat", opencv.IMDecodeIntoMat, true)
	engine.RegisterMethod("opencv.arcLength", "OpenCV.ArcLength", opencv.ArcLength, true)
	engine.RegisterMethod("opencv.approxPolyDP", "OpenCV.ApproxPolyDP", opencv.ApproxPolyDP, true)
	engine.RegisterMethod("opencv.convexHull", "OpenCV.ConvexHull", opencv.ConvexHull, true)
	engine.RegisterMethod("opencv.convexityDefects", "OpenCV.ConvexityDefects", opencv.ConvexityDefects, true)
	engine.RegisterMethod("opencv.cvtColor", "OpenCV.CvtColor", opencv.CvtColor, true)
	engine.RegisterMethod("opencv.demosaicing", "OpenCV.Demosaicing", opencv.Demosaicing, true)
	engine.RegisterMethod("opencv.equalizeHist", "OpenCV.EqualizeHist", opencv.EqualizeHist, true)
	engine.RegisterMethod("opencv.calcHist", "OpenCV.CalcHist", opencv.CalcHist, true)
	engine.RegisterMethod("opencv.calcBackProject", "OpenCV.CalcBackProject", opencv.CalcBackProject, true)
	engine.RegisterMethod("opencv.compareHist", "OpenCV.CompareHist", opencv.CompareHist, true)
	engine.RegisterMethod("opencv.eMD", "OpenCV.EMD", opencv.EMD, true)
	engine.RegisterMethod("opencv.clipLine", "OpenCV.ClipLine", opencv.ClipLine, true)
	engine.RegisterMethod("opencv.bilateralFilter", "OpenCV.BilateralFilter", opencv.BilateralFilter, true)
	engine.RegisterMethod("opencv.blur", "OpenCV.Blur", opencv.Blur, true)
	engine.RegisterMethod("opencv.boxFilter", "OpenCV.BoxFilter", opencv.BoxFilter, true)
	engine.RegisterMethod("opencv.sqBoxFilter", "OpenCV.SqBoxFilter", opencv.SqBoxFilter, true)
	engine.RegisterMethod("opencv.dilate", "OpenCV.Dilate", opencv.Dilate, true)
	engine.RegisterMethod("opencv.dilateWithParams", "OpenCV.DilateWithParams", opencv.DilateWithParams, true)
	engine.RegisterMethod("opencv.distanceTransform", "OpenCV.DistanceTransform", opencv.DistanceTransform, true)
	engine.RegisterMethod("opencv.erode", "OpenCV.Erode", opencv.Erode, true)
	engine.RegisterMethod("opencv.erodeWithParams", "OpenCV.ErodeWithParams", opencv.ErodeWithParams, true)
	engine.RegisterMethod("opencv.erodeWithParamsAndBorderValue", "OpenCV.ErodeWithParamsAndBorderValue", opencv.ErodeWithParamsAndBorderValue, true)
	engine.RegisterMethod("opencv.boundingRect", "OpenCV.BoundingRect", opencv.BoundingRect, true)
	engine.RegisterMethod("opencv.boxPoints", "OpenCV.BoxPoints", opencv.BoxPoints, true)
	engine.RegisterMethod("opencv.boxPoints2f", "OpenCV.BoxPoints2f", opencv.BoxPoints2f, true)
	engine.RegisterMethod("opencv.contourArea", "OpenCV.ContourArea", opencv.ContourArea, true)
	engine.RegisterMethod("opencv.minAreaRect", "OpenCV.MinAreaRect", opencv.MinAreaRect, true)
	engine.RegisterMethod("opencv.minAreaRect2f", "OpenCV.MinAreaRect2f", opencv.MinAreaRect2f, true)
	engine.RegisterMethod("opencv.fitEllipse", "OpenCV.FitEllipse", opencv.FitEllipse, true)
	engine.RegisterMethod("opencv.minEnclosingCircle", "OpenCV.MinEnclosingCircle", opencv.MinEnclosingCircle, true)
	engine.RegisterMethod("opencv.findContours", "OpenCV.FindContours", opencv.FindContours, true)
	engine.RegisterMethod("opencv.findContoursWithParams", "OpenCV.FindContoursWithParams", opencv.FindContoursWithParams, true)
	engine.RegisterMethod("opencv.pointPolygonTest", "OpenCV.PointPolygonTest", opencv.PointPolygonTest, true)
	engine.RegisterMethod("opencv.connectedComponents", "OpenCV.ConnectedComponents", opencv.ConnectedComponents, true)
	engine.RegisterMethod("opencv.connectedComponentsWithParams", "OpenCV.ConnectedComponentsWithParams", opencv.ConnectedComponentsWithParams, true)
	engine.RegisterMethod("opencv.connectedComponentsWithStats", "OpenCV.ConnectedComponentsWithStats", opencv.ConnectedComponentsWithStats, true)
	engine.RegisterMethod("opencv.connectedComponentsWithStatsWithParams", "OpenCV.ConnectedComponentsWithStatsWithParams", opencv.ConnectedComponentsWithStatsWithParams, true)
	engine.RegisterMethod("opencv.matchTemplate", "OpenCV.MatchTemplate", opencv.MatchTemplate, true)
	engine.RegisterMethod("opencv.moments", "OpenCV.Moments", opencv.Moments, true)
	engine.RegisterMethod("opencv.pyrDown", "OpenCV.PyrDown", opencv.PyrDown, true)
	engine.RegisterMethod("opencv.pyrUp", "OpenCV.PyrUp", opencv.PyrUp, true)
	engine.RegisterMethod("opencv.morphologyDefaultBorderValue", "OpenCV.MorphologyDefaultBorderValue", opencv.MorphologyDefaultBorderValue, true)
	engine.RegisterMethod("opencv.morphologyEx", "OpenCV.MorphologyEx", opencv.MorphologyEx, true)
	engine.RegisterMethod("opencv.morphologyExWithParams", "OpenCV.MorphologyExWithParams", opencv.MorphologyExWithParams, true)
	engine.RegisterMethod("opencv.getStructuringElement", "OpenCV.GetStructuringElement", opencv.GetStructuringElement, true)
	engine.RegisterMethod("opencv.gaussianBlur", "OpenCV.GaussianBlur", opencv.GaussianBlur, true)
	engine.RegisterMethod("opencv.getGaussianKernel", "OpenCV.GetGaussianKernel", opencv.GetGaussianKernel, true)
	engine.RegisterMethod("opencv.getGaussianKernelWithParams", "OpenCV.GetGaussianKernelWithParams", opencv.GetGaussianKernelWithParams, true)
	engine.RegisterMethod("opencv.sobel", "OpenCV.Sobel", opencv.Sobel, true)
	engine.RegisterMethod("opencv.spatialGradient", "OpenCV.SpatialGradient", opencv.SpatialGradient, true)
	engine.RegisterMethod("opencv.laplacian", "OpenCV.Laplacian", opencv.Laplacian, true)
	engine.RegisterMethod("opencv.scharr", "OpenCV.Scharr", opencv.Scharr, true)
	engine.RegisterMethod("opencv.medianBlur", "OpenCV.MedianBlur", opencv.MedianBlur, true)
	engine.RegisterMethod("opencv.canny", "OpenCV.Canny", opencv.Canny, true)
	engine.RegisterMethod("opencv.cornerSubPix", "OpenCV.CornerSubPix", opencv.CornerSubPix, true)
	engine.RegisterMethod("opencv.goodFeaturesToTrack", "OpenCV.GoodFeaturesToTrack", opencv.GoodFeaturesToTrack, true)
	engine.RegisterMethod("opencv.grabCut", "OpenCV.GrabCut", opencv.GrabCut, true)
	engine.RegisterMethod("opencv.houghCircles", "OpenCV.HoughCircles", opencv.HoughCircles, true)
	engine.RegisterMethod("opencv.houghCirclesWithParams", "OpenCV.HoughCirclesWithParams", opencv.HoughCirclesWithParams, true)
	engine.RegisterMethod("opencv.houghLines", "OpenCV.HoughLines", opencv.HoughLines, true)
	engine.RegisterMethod("opencv.houghLinesP", "OpenCV.HoughLinesP", opencv.HoughLinesP, true)
	engine.RegisterMethod("opencv.houghLinesPWithParams", "OpenCV.HoughLinesPWithParams", opencv.HoughLinesPWithParams, true)
	engine.RegisterMethod("opencv.houghLinesPointSet", "OpenCV.HoughLinesPointSet", opencv.HoughLinesPointSet, true)
	engine.RegisterMethod("opencv.integral", "OpenCV.Integral", opencv.Integral, true)
	engine.RegisterMethod("opencv.threshold", "OpenCV.Threshold", opencv.Threshold, true)
	engine.RegisterMethod("opencv.adaptiveThreshold", "OpenCV.AdaptiveThreshold", opencv.AdaptiveThreshold, true)
	engine.RegisterMethod("opencv.arrowedLine", "OpenCV.ArrowedLine", opencv.ArrowedLine, true)
	engine.RegisterMethod("opencv.circle", "OpenCV.Circle", opencv.Circle, true)
	engine.RegisterMethod("opencv.circleWithParams", "OpenCV.CircleWithParams", opencv.CircleWithParams, true)
	engine.RegisterMethod("opencv.ellipse", "OpenCV.Ellipse", opencv.Ellipse, true)
	engine.RegisterMethod("opencv.ellipseWithParams", "OpenCV.EllipseWithParams", opencv.EllipseWithParams, true)
	engine.RegisterMethod("opencv.line", "OpenCV.Line", opencv.Line, true)
	engine.RegisterMethod("opencv.rectangle", "OpenCV.Rectangle", opencv.Rectangle, true)
	engine.RegisterMethod("opencv.rectangleWithParams", "OpenCV.RectangleWithParams", opencv.RectangleWithParams, true)
	engine.RegisterMethod("opencv.fillPoly", "OpenCV.FillPoly", opencv.FillPoly, true)
	engine.RegisterMethod("opencv.fillPolyWithParams", "OpenCV.FillPolyWithParams", opencv.FillPolyWithParams, true)
	engine.RegisterMethod("opencv.polylines", "OpenCV.Polylines", opencv.Polylines, true)
	engine.RegisterMethod("opencv.getTextSize", "OpenCV.GetTextSize", opencv.GetTextSize, true)
	engine.RegisterMethod("opencv.getTextSizeWithBaseline", "OpenCV.GetTextSizeWithBaseline", opencv.GetTextSizeWithBaseline, true)
	engine.RegisterMethod("opencv.putText", "OpenCV.PutText", opencv.PutText, true)
	engine.RegisterMethod("opencv.putTextWithParams", "OpenCV.PutTextWithParams", opencv.PutTextWithParams, true)
	engine.RegisterMethod("opencv.resize", "OpenCV.Resize", opencv.Resize, true)
	engine.RegisterMethod("opencv.getRectSubPix", "OpenCV.GetRectSubPix", opencv.GetRectSubPix, true)
	engine.RegisterMethod("opencv.getRotationMatrix2D", "OpenCV.GetRotationMatrix2D", opencv.GetRotationMatrix2D, true)
	engine.RegisterMethod("opencv.warpAffine", "OpenCV.WarpAffine", opencv.WarpAffine, true)
	engine.RegisterMethod("opencv.warpAffineWithParams", "OpenCV.WarpAffineWithParams", opencv.WarpAffineWithParams, true)
	engine.RegisterMethod("opencv.warpPerspective", "OpenCV.WarpPerspective", opencv.WarpPerspective, true)
	engine.RegisterMethod("opencv.warpPerspectiveWithParams", "OpenCV.WarpPerspectiveWithParams", opencv.WarpPerspectiveWithParams, true)
	engine.RegisterMethod("opencv.watershed", "OpenCV.Watershed", opencv.Watershed, true)
	engine.RegisterMethod("opencv.applyColorMap", "OpenCV.ApplyColorMap", opencv.ApplyColorMap, true)
	engine.RegisterMethod("opencv.applyCustomColorMap", "OpenCV.ApplyCustomColorMap", opencv.ApplyCustomColorMap, true)
	engine.RegisterMethod("opencv.getPerspectiveTransform", "OpenCV.GetPerspectiveTransform", opencv.GetPerspectiveTransform, true)
	engine.RegisterMethod("opencv.getPerspectiveTransform2f", "OpenCV.GetPerspectiveTransform2f", opencv.GetPerspectiveTransform2f, true)
	engine.RegisterMethod("opencv.getAffineTransform", "OpenCV.GetAffineTransform", opencv.GetAffineTransform, true)
	engine.RegisterMethod("opencv.getAffineTransform2f", "OpenCV.GetAffineTransform2f", opencv.GetAffineTransform2f, true)
	engine.RegisterMethod("opencv.drawContours", "OpenCV.DrawContours", opencv.DrawContours, true)
	engine.RegisterMethod("opencv.drawContoursWithParams", "OpenCV.DrawContoursWithParams", opencv.DrawContoursWithParams, true)
	engine.RegisterMethod("opencv.remap", "OpenCV.Remap", opencv.Remap, true)
	engine.RegisterMethod("opencv.filter2D", "OpenCV.Filter2D", opencv.Filter2D, true)
	engine.RegisterMethod("opencv.sepFilter2D", "OpenCV.SepFilter2D", opencv.SepFilter2D, true)
	engine.RegisterMethod("opencv.logPolar", "OpenCV.LogPolar", opencv.LogPolar, true)
	engine.RegisterMethod("opencv.linearPolar", "OpenCV.LinearPolar", opencv.LinearPolar, true)
	engine.RegisterMethod("opencv.fitLine", "OpenCV.FitLine", opencv.FitLine, true)
	engine.RegisterMethod("opencv.matchShapes", "OpenCV.MatchShapes", opencv.MatchShapes, true)
	engine.RegisterMethod("opencv.newCLAHE", "OpenCV.NewCLAHE", opencv.NewCLAHE, true)
	engine.RegisterMethod("opencv.newCLAHEWithParams", "OpenCV.NewCLAHEWithParams", opencv.NewCLAHEWithParams, true)
	engine.RegisterMethod("opencv.invertAffineTransform", "OpenCV.InvertAffineTransform", opencv.InvertAffineTransform, true)
	engine.RegisterMethod("opencv.phaseCorrelate", "OpenCV.PhaseCorrelate", opencv.PhaseCorrelate, true)
	engine.RegisterMethod("opencv.createHanningWindow", "OpenCV.CreateHanningWindow", opencv.CreateHanningWindow, true)
	engine.RegisterMethod("opencv.imageToMatRGBA", "OpenCV.ImageToMatRGBA", opencv.ImageToMatRGBA, true)
	engine.RegisterMethod("opencv.imageToMatRGB", "OpenCV.ImageToMatRGB", opencv.ImageToMatRGB, true)
	engine.RegisterMethod("opencv.imageGrayToMatGray", "OpenCV.ImageGrayToMatGray", opencv.ImageGrayToMatGray, true)
	engine.RegisterMethod("opencv.accumulate", "OpenCV.Accumulate", opencv.Accumulate, true)
	engine.RegisterMethod("opencv.accumulateWithMask", "OpenCV.AccumulateWithMask", opencv.AccumulateWithMask, true)
	engine.RegisterMethod("opencv.accumulateSquare", "OpenCV.AccumulateSquare", opencv.AccumulateSquare, true)
	engine.RegisterMethod("opencv.accumulateSquareWithMask", "OpenCV.AccumulateSquareWithMask", opencv.AccumulateSquareWithMask, true)
	engine.RegisterMethod("opencv.accumulateProduct", "OpenCV.AccumulateProduct", opencv.AccumulateProduct, true)
	engine.RegisterMethod("opencv.accumulateProductWithMask", "OpenCV.AccumulateProductWithMask", opencv.AccumulateProductWithMask, true)
	engine.RegisterMethod("opencv.accumulatedWeighted", "OpenCV.AccumulatedWeighted", opencv.AccumulatedWeighted, true)
	engine.RegisterMethod("opencv.accumulatedWeightedWithMask", "OpenCV.AccumulatedWeightedWithMask", opencv.AccumulatedWeightedWithMask, true)
	engine.RegisterMethod("opencv.findImage", "OpenCV.FindImage", opencv.FindImage, true)
	engine.RegisterMethod("opencv.findImageAll", "OpenCV.FindImageAll", opencv.FindImageAll, true)
	engine.RegisterMethod("opencv.close", "OpenCV Mat.Close", (*opencv.Mat).Close, true)
	engine.RegisterMethod("opencv.fromPtr", "OpenCV Mat.FromPtr", (*opencv.Mat).FromPtr, true)
	engine.RegisterMethod("opencv.ptr", "OpenCV Mat.Ptr", (*opencv.Mat).Ptr, true)
	engine.RegisterMethod("opencv.empty", "OpenCV Mat.Empty", (*opencv.Mat).Empty, true)
	engine.RegisterMethod("opencv.closed", "OpenCV Mat.Closed", (*opencv.Mat).Closed, true)
	engine.RegisterMethod("opencv.isContinuous", "OpenCV Mat.IsContinuous", (*opencv.Mat).IsContinuous, true)
	engine.RegisterMethod("opencv.inv", "OpenCV Mat.Inv", (*opencv.Mat).Inv, true)
	engine.RegisterMethod("opencv.col", "OpenCV Mat.Col", (*opencv.Mat).Col, true)
	engine.RegisterMethod("opencv.row", "OpenCV Mat.Row", (*opencv.Mat).Row, true)
	engine.RegisterMethod("opencv.clone", "OpenCV Mat.Clone", (*opencv.Mat).Clone, true)
	engine.RegisterMethod("opencv.copyTo", "OpenCV Mat.CopyTo", (*opencv.Mat).CopyTo, true)
	engine.RegisterMethod("opencv.copyToWithMask", "OpenCV Mat.CopyToWithMask", (*opencv.Mat).CopyToWithMask, true)
	engine.RegisterMethod("opencv.convertTo", "OpenCV Mat.ConvertTo", (*opencv.Mat).ConvertTo, true)
	engine.RegisterMethod("opencv.convertToWithParams", "OpenCV Mat.ConvertToWithParams", (*opencv.Mat).ConvertToWithParams, true)
	engine.RegisterMethod("opencv.total", "OpenCV Mat.Total", (*opencv.Mat).Total, true)
	engine.RegisterMethod("opencv.size", "OpenCV Mat.Size", (*opencv.Mat).Size, true)
	engine.RegisterMethod("opencv.toBytes", "OpenCV Mat.ToBytes", (*opencv.Mat).ToBytes, true)
	engine.RegisterMethod("opencv.dataPtrUint8", "OpenCV Mat.DataPtrUint8", (*opencv.Mat).DataPtrUint8, true)
	engine.RegisterMethod("opencv.dataPtrInt8", "OpenCV Mat.DataPtrInt8", (*opencv.Mat).DataPtrInt8, true)
	engine.RegisterMethod("opencv.dataPtrUint16", "OpenCV Mat.DataPtrUint16", (*opencv.Mat).DataPtrUint16, true)
	engine.RegisterMethod("opencv.dataPtrInt16", "OpenCV Mat.DataPtrInt16", (*opencv.Mat).DataPtrInt16, true)
	engine.RegisterMethod("opencv.dataPtrFloat32", "OpenCV Mat.DataPtrFloat32", (*opencv.Mat).DataPtrFloat32, true)
	engine.RegisterMethod("opencv.dataPtrFloat64", "OpenCV Mat.DataPtrFloat64", (*opencv.Mat).DataPtrFloat64, true)
	engine.RegisterMethod("opencv.region", "OpenCV Mat.Region", (*opencv.Mat).Region, true)
	engine.RegisterMethod("opencv.reshape", "OpenCV Mat.Reshape", (*opencv.Mat).Reshape, true)
	engine.RegisterMethod("opencv.convertFp16", "OpenCV Mat.ConvertFp16", (*opencv.Mat).ConvertFp16, true)
	engine.RegisterMethod("opencv.mean", "OpenCV Mat.Mean", (*opencv.Mat).Mean, true)
	engine.RegisterMethod("opencv.meanWithMask", "OpenCV Mat.MeanWithMask", (*opencv.Mat).MeanWithMask, true)
	engine.RegisterMethod("opencv.sqrt", "OpenCV Mat.Sqrt", (*opencv.Mat).Sqrt, true)
	engine.RegisterMethod("opencv.sum", "OpenCV Mat.Sum", (*opencv.Mat).Sum, true)
	engine.RegisterMethod("opencv.patchNaNs", "OpenCV Mat.PatchNaNs", (*opencv.Mat).PatchNaNs, true)
	engine.RegisterMethod("opencv.rows", "OpenCV Mat.Rows", (*opencv.Mat).Rows, true)
	engine.RegisterMethod("opencv.cols", "OpenCV Mat.Cols", (*opencv.Mat).Cols, true)
	engine.RegisterMethod("opencv.channels", "OpenCV Mat.Channels", (*opencv.Mat).Channels, true)
	engine.RegisterMethod("opencv.type", "OpenCV Mat.Type", (*opencv.Mat).Type, true)
	engine.RegisterMethod("opencv.step", "OpenCV Mat.Step", (*opencv.Mat).Step, true)
	engine.RegisterMethod("opencv.elemSize", "OpenCV Mat.ElemSize", (*opencv.Mat).ElemSize, true)
	engine.RegisterMethod("opencv.getUCharAt", "OpenCV Mat.GetUCharAt", (*opencv.Mat).GetUCharAt, true)
	engine.RegisterMethod("opencv.getUCharAt3", "OpenCV Mat.GetUCharAt3", (*opencv.Mat).GetUCharAt3, true)
	engine.RegisterMethod("opencv.getSCharAt", "OpenCV Mat.GetSCharAt", (*opencv.Mat).GetSCharAt, true)
	engine.RegisterMethod("opencv.getSCharAt3", "OpenCV Mat.GetSCharAt3", (*opencv.Mat).GetSCharAt3, true)
	engine.RegisterMethod("opencv.getShortAt", "OpenCV Mat.GetShortAt", (*opencv.Mat).GetShortAt, true)
	engine.RegisterMethod("opencv.getShortAt3", "OpenCV Mat.GetShortAt3", (*opencv.Mat).GetShortAt3, true)
	engine.RegisterMethod("opencv.getIntAt", "OpenCV Mat.GetIntAt", (*opencv.Mat).GetIntAt, true)
	engine.RegisterMethod("opencv.getIntAt3", "OpenCV Mat.GetIntAt3", (*opencv.Mat).GetIntAt3, true)
	engine.RegisterMethod("opencv.getFloatAt", "OpenCV Mat.GetFloatAt", (*opencv.Mat).GetFloatAt, true)
	engine.RegisterMethod("opencv.getFloatAt3", "OpenCV Mat.GetFloatAt3", (*opencv.Mat).GetFloatAt3, true)
	engine.RegisterMethod("opencv.getDoubleAt", "OpenCV Mat.GetDoubleAt", (*opencv.Mat).GetDoubleAt, true)
	engine.RegisterMethod("opencv.getDoubleAt3", "OpenCV Mat.GetDoubleAt3", (*opencv.Mat).GetDoubleAt3, true)
	engine.RegisterMethod("opencv.setTo", "OpenCV Mat.SetTo", (*opencv.Mat).SetTo, true)
	engine.RegisterMethod("opencv.setUCharAt", "OpenCV Mat.SetUCharAt", (*opencv.Mat).SetUCharAt, true)
	engine.RegisterMethod("opencv.setUCharAt3", "OpenCV Mat.SetUCharAt3", (*opencv.Mat).SetUCharAt3, true)
	engine.RegisterMethod("opencv.setSCharAt", "OpenCV Mat.SetSCharAt", (*opencv.Mat).SetSCharAt, true)
	engine.RegisterMethod("opencv.setSCharAt3", "OpenCV Mat.SetSCharAt3", (*opencv.Mat).SetSCharAt3, true)
	engine.RegisterMethod("opencv.setShortAt", "OpenCV Mat.SetShortAt", (*opencv.Mat).SetShortAt, true)
	engine.RegisterMethod("opencv.setShortAt3", "OpenCV Mat.SetShortAt3", (*opencv.Mat).SetShortAt3, true)
	engine.RegisterMethod("opencv.setIntAt", "OpenCV Mat.SetIntAt", (*opencv.Mat).SetIntAt, true)
	engine.RegisterMethod("opencv.setIntAt3", "OpenCV Mat.SetIntAt3", (*opencv.Mat).SetIntAt3, true)
	engine.RegisterMethod("opencv.setFloatAt", "OpenCV Mat.SetFloatAt", (*opencv.Mat).SetFloatAt, true)
	engine.RegisterMethod("opencv.setFloatAt3", "OpenCV Mat.SetFloatAt3", (*opencv.Mat).SetFloatAt3, true)
	engine.RegisterMethod("opencv.setDoubleAt", "OpenCV Mat.SetDoubleAt", (*opencv.Mat).SetDoubleAt, true)
	engine.RegisterMethod("opencv.setDoubleAt3", "OpenCV Mat.SetDoubleAt3", (*opencv.Mat).SetDoubleAt3, true)
	engine.RegisterMethod("opencv.addUChar", "OpenCV Mat.AddUChar", (*opencv.Mat).AddUChar, true)
	engine.RegisterMethod("opencv.subtractUChar", "OpenCV Mat.SubtractUChar", (*opencv.Mat).SubtractUChar, true)
	engine.RegisterMethod("opencv.multiplyUChar", "OpenCV Mat.MultiplyUChar", (*opencv.Mat).MultiplyUChar, true)
	engine.RegisterMethod("opencv.divideUChar", "OpenCV Mat.DivideUChar", (*opencv.Mat).DivideUChar, true)
	engine.RegisterMethod("opencv.addFloat", "OpenCV Mat.AddFloat", (*opencv.Mat).AddFloat, true)
	engine.RegisterMethod("opencv.subtractFloat", "OpenCV Mat.SubtractFloat", (*opencv.Mat).SubtractFloat, true)
	engine.RegisterMethod("opencv.multiplyFloat", "OpenCV Mat.MultiplyFloat", (*opencv.Mat).MultiplyFloat, true)
	engine.RegisterMethod("opencv.divideFloat", "OpenCV Mat.DivideFloat", (*opencv.Mat).DivideFloat, true)
	engine.RegisterMethod("opencv.multiplyMatrix", "OpenCV Mat.MultiplyMatrix", (*opencv.Mat).MultiplyMatrix, true)
	engine.RegisterMethod("opencv.t", "OpenCV Mat.T", (*opencv.Mat).T, true)
	engine.RegisterMethod("opencv.getVecbAt", "OpenCV Mat.GetVecbAt", (*opencv.Mat).GetVecbAt, true)
	engine.RegisterMethod("opencv.getVecfAt", "OpenCV Mat.GetVecfAt", (*opencv.Mat).GetVecfAt, true)
	engine.RegisterMethod("opencv.getVecdAt", "OpenCV Mat.GetVecdAt", (*opencv.Mat).GetVecdAt, true)
	engine.RegisterMethod("opencv.getVeciAt", "OpenCV Mat.GetVeciAt", (*opencv.Mat).GetVeciAt, true)
	engine.RegisterMethod("opencv.rowRange", "OpenCV Mat.RowRange", (*opencv.Mat).RowRange, true)
	engine.RegisterMethod("opencv.colRange", "OpenCV Mat.ColRange", (*opencv.Mat).ColRange, true)
	engine.RegisterMethod("opencv.toImage", "OpenCV Mat.ToImage", (*opencv.Mat).ToImage, true)
	engine.RegisterMethod("opencv.toImageYUV", "OpenCV Mat.ToImageYUV", (*opencv.Mat).ToImageYUV, true)
	engine.RegisterMethod("opencv.toImageYUVWithParams", "OpenCV Mat.ToImageYUVWithParams", (*opencv.Mat).ToImageYUVWithParams, true)
	engine.RegisterMethod("opencv.isNil", "OpenCV PointVector.IsNil", opencv.PointVector.IsNil, true)
	engine.RegisterMethod("opencv.at", "OpenCV PointVector.At", opencv.PointVector.At, true)
	engine.RegisterMethod("opencv.append", "OpenCV PointVector.Append", opencv.PointVector.Append, true)
	engine.RegisterMethod("opencv.toPoints", "OpenCV PointVector.ToPoints", opencv.PointVector.ToPoints, true)
	engine.RegisterMethod("opencv.p", "OpenCV PointsVector.P", opencv.PointsVector.P, true)
	engine.RegisterMethod("opencv.fill", "OpenCV RNG.Fill", (*opencv.RNG).Fill, true)
	engine.RegisterMethod("opencv.gaussian", "OpenCV RNG.Gaussian", (*opencv.RNG).Gaussian, true)
	engine.RegisterMethod("opencv.next", "OpenCV RNG.Next", (*opencv.RNG).Next, true)
	engine.RegisterMethod("opencv.getBytes", "OpenCV NativeByteBuffer.GetBytes", (*opencv.NativeByteBuffer).GetBytes, true)
	engine.RegisterMethod("opencv.len", "OpenCV NativeByteBuffer.Len", (*opencv.NativeByteBuffer).Len, true)
	engine.RegisterMethod("opencv.apply", "OpenCV CLAHE.Apply", (*opencv.CLAHE).Apply, true)

	return nil
}

func bindOpenCVLuaFunc(fn reflect.Value) lua.LGFunction {
	fnType := fn.Type()
	return func(L *lua.LState) (resultCount int) {
		defer func() {
			if recover() != nil {
				L.Push(lua.LNil)
				resultCount = 1
			}
		}()
		args := buildOpenCVLuaArgs(L, fnType)
		if fnType.IsVariadic() {
			return pushOpenCVLuaResults(L, fn.CallSlice(args))
		}
		return pushOpenCVLuaResults(L, fn.Call(args))
	}
}

func bindOpenCVLuaMethod(value reflect.Value, methodName string) lua.LGFunction {
	return func(L *lua.LState) (resultCount int) {
		defer func() {
			if recover() != nil {
				L.Push(lua.LNil)
				resultCount = 1
			}
		}()
		method := value.MethodByName(methodName)
		if !method.IsValid() {
			L.Push(lua.LNil)
			return 1
		}
		methodType := method.Type()
		args := buildOpenCVLuaArgs(L, methodType)
		if methodType.IsVariadic() {
			return pushOpenCVLuaResults(L, method.CallSlice(args))
		}
		return pushOpenCVLuaResults(L, method.Call(args))
	}
}

func buildOpenCVLuaArgs(L *lua.LState, fnType reflect.Type) []reflect.Value {
	if !fnType.IsVariadic() {
		args := make([]reflect.Value, fnType.NumIn())
		for i := 0; i < fnType.NumIn(); i++ {
			args[i] = luaToReflect(L, L.Get(i+1), fnType.In(i))
		}
		return args
	}
	fixedCount := fnType.NumIn() - 1
	args := make([]reflect.Value, fnType.NumIn())
	for i := 0; i < fixedCount; i++ {
		args[i] = luaToReflect(L, L.Get(i+1), fnType.In(i))
	}
	variadicType := fnType.In(fnType.NumIn() - 1)
	variadic := reflect.MakeSlice(variadicType, 0, maxInt(L.GetTop()-fixedCount, 0))
	for i := fixedCount; i < L.GetTop(); i++ {
		variadic = reflect.Append(variadic, luaToReflect(L, L.Get(i+1), variadicType.Elem()))
	}
	args[fnType.NumIn()-1] = variadic
	return args
}

func luaToReflect(L *lua.LState, value lua.LValue, target reflect.Type) reflect.Value {
	if target.Kind() == reflect.Interface {
		return reflect.ValueOf(luaToInterface(value))
	}
	if target.Kind() == reflect.Ptr {
		if target.Elem().Kind() == reflect.Slice {
			if target.Elem().Elem().Kind() != reflect.Uint8 {
				ptr := reflect.New(target.Elem())
				ptr.Elem().Set(luaToReflect(L, value, target.Elem()))
				return ptr
			}
			bytes := []byte(lua.LVAsString(value))
			ptr := reflect.New(target.Elem())
			ptr.Elem().Set(reflect.ValueOf(bytes))
			return ptr
		}
		if unwrapped, ok := unwrapLuaGoValue(value); ok {
			if unwrapped.Type().AssignableTo(target) {
				return unwrapped
			}
			if unwrapped.Type().ConvertibleTo(target) {
				return unwrapped.Convert(target)
			}
		}
		if target.Elem().Kind() == reflect.Struct {
			ptr := reflect.New(target.Elem())
			fillLuaStruct(ptr.Elem(), value)
			return ptr
		}
		return reflect.Zero(target)
	}
	if unwrapped, ok := unwrapLuaGoValue(value); ok {
		if unwrapped.Type().AssignableTo(target) {
			return unwrapped
		}
		if unwrapped.Type().ConvertibleTo(target) {
			return unwrapped.Convert(target)
		}
	}
	switch target.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(lua.LVAsBool(value)).Convert(target)
	case reflect.String:
		return reflect.ValueOf(lua.LVAsString(value)).Convert(target)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(int64(lua.LVAsNumber(value))).Convert(target)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return reflect.ValueOf(uint64(lua.LVAsNumber(value))).Convert(target)
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(float64(lua.LVAsNumber(value))).Convert(target)
	case reflect.Slice:
		if target.Elem().Kind() == reflect.Uint8 {
			return reflect.ValueOf([]byte(lua.LVAsString(value))).Convert(target)
		}
		table, ok := value.(*lua.LTable)
		if !ok {
			return reflect.Zero(target)
		}
		result := reflect.MakeSlice(target, 0, table.Len())
		for i := 1; i <= table.Len(); i++ {
			result = reflect.Append(result, luaToReflect(L, table.RawGetInt(i), target.Elem()))
		}
		return result
	case reflect.Array:
		table, ok := value.(*lua.LTable)
		if !ok {
			return reflect.Zero(target)
		}
		result := reflect.New(target).Elem()
		for i := 0; i < target.Len(); i++ {
			result.Index(i).Set(luaToReflect(L, table.RawGetInt(i+1), target.Elem()))
		}
		return result
	case reflect.Map:
		table, ok := value.(*lua.LTable)
		if !ok {
			return reflect.Zero(target)
		}
		result := reflect.MakeMap(target)
		table.ForEach(func(key lua.LValue, item lua.LValue) {
			result.SetMapIndex(luaToReflect(L, key, target.Key()), luaToReflect(L, item, target.Elem()))
		})
		return result
	case reflect.Struct:
		result := reflect.New(target).Elem()
		fillLuaStruct(result, value)
		return result
	}
	return reflect.Zero(target)
}

func fillLuaStruct(result reflect.Value, value lua.LValue) {
	table, ok := value.(*lua.LTable)
	if !ok {
		return
	}
	resultType := result.Type()
	for i := 0; i < result.NumField(); i++ {
		field := result.Field(i)
		if !field.CanSet() {
			continue
		}
		fieldInfo := resultType.Field(i)
		fieldValue := table.RawGetString(fieldInfo.Name)
		if fieldValue == lua.LNil {
			fieldValue = table.RawGetString(lowerFirst(fieldInfo.Name))
		}
		if fieldValue == lua.LNil {
			continue
		}
		field.Set(luaToReflect(nil, fieldValue, field.Type()))
	}
}

func lowerFirst(value string) string {
	if value == "" {
		return ""
	}
	first := value[0]
	if first >= 'A' && first <= 'Z' {
		first += 'a' - 'A'
	}
	return string(first) + value[1:]
}

func unwrapLuaGoValue(value lua.LValue) (reflect.Value, bool) {
	if ud, ok := value.(*lua.LUserData); ok && ud.Value != nil {
		return reflect.ValueOf(ud.Value), true
	}
	if table, ok := value.(*lua.LTable); ok {
		if ud, ok := table.RawGetString("__go").(*lua.LUserData); ok && ud.Value != nil {
			return reflect.ValueOf(ud.Value), true
		}
	}
	return reflect.Value{}, false
}

func luaToInterface(value lua.LValue) interface{} {
	if unwrapped, ok := unwrapLuaGoValue(value); ok {
		return unwrapped.Interface()
	}
	switch v := value.(type) {
	case lua.LBool:
		return bool(v)
	case lua.LNumber:
		return float64(v)
	case lua.LString:
		return string(v)
	default:
		return nil
	}
}

func pushOpenCVLuaResults(L *lua.LState, results []reflect.Value) int {
	if len(results) == 0 {
		return 0
	}
	if len(results) == 1 {
		L.Push(reflectToLua(L, results[0]))
		return 1
	}
	obj := L.NewTable()
	for i, result := range results {
		obj.RawSetString("r"+strconvItoa(i), reflectToLua(L, result))
	}
	if len(results) == 2 && isReflectInt(results[0]) && isReflectInt(results[1]) {
		obj.RawSetString("x", reflectToLua(L, results[0]))
		obj.RawSetString("y", reflectToLua(L, results[1]))
	}
	L.Push(obj)
	return 1
}

func reflectToLua(L *lua.LState, value reflect.Value) lua.LValue {
	if !value.IsValid() {
		return lua.LNil
	}
	if value.Kind() == reflect.Ptr && value.IsNil() {
		return lua.LNil
	}
	switch value.Kind() {
	case reflect.Bool:
		return lua.LBool(value.Bool())
	case reflect.String:
		return lua.LString(value.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return lua.LNumber(value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return lua.LNumber(value.Uint())
	case reflect.Float32, reflect.Float64:
		return lua.LNumber(value.Float())
	case reflect.Slice, reflect.Array:
		if value.Type().Elem().Kind() == reflect.Uint8 {
			return lua.LString(string(value.Bytes()))
		}
		arr := L.NewTable()
		for i := 0; i < value.Len(); i++ {
			arr.RawSetInt(i+1, reflectToLua(L, value.Index(i)))
		}
		return arr
	case reflect.Map:
		obj := L.NewTable()
		iter := value.MapRange()
		for iter.Next() {
			obj.RawSet(reflectToLua(L, iter.Key()), reflectToLua(L, iter.Value()))
		}
		return obj
	}
	return wrapOpenCVLuaObject(L, value)
}

func wrapOpenCVLuaObject(L *lua.LState, value reflect.Value) lua.LValue {
	if value.Kind() == reflect.Ptr && value.IsNil() {
		return lua.LNil
	}
	obj := L.NewTable()
	ud := L.NewUserData()
	ud.Value = value.Interface()
	obj.RawSetString("__go", ud)
	typ := value.Type()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	for _, entry := range openCVMethods[typ.Name()] {
		methodName := entry.goMethod
		obj.RawSetString(entry.name, L.NewFunction(bindOpenCVLuaMethod(value, methodName)))
	}
	return obj
}

func isReflectInt(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func maxInt(left int, right int) int {
	if left > right {
		return left
	}
	return right
}

func strconvItoa(value int) string {
	if value == 0 {
		return "0"
	}
	digits := [20]byte{}
	i := len(digits)
	for value > 0 {
		i--
		digits[i] = byte('0' + value%10)
		value /= 10
	}
	return string(digits[i:])
}
