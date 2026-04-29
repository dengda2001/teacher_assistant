import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Progress } from "@/components/ui/progress"
import {
  Users,
  GraduationCap,
  BookOpen,
  Award,
  ArrowUpRight,
  ArrowDownRight,
  Calendar,
  Clock
} from "lucide-react"

// Mock data
const stats = [
  {
    title: "学生总数",
    value: "48",
    change: "+3",
    trend: "up",
    icon: Users,
    description: "本学期新增3名学生"
  },
  {
    title: "平均分",
    value: "82.5",
    change: "+2.3",
    trend: "up",
    icon: GraduationCap,
    description: "较上月提升2.3分"
  },
  {
    title: "及格率",
    value: "92%",
    change: "+5%",
    trend: "up",
    icon: BookOpen,
    description: "较上月提升5%"
  },
  {
    title: "优秀率",
    value: "35%",
    change: "-2%",
    trend: "down",
    icon: Award,
    description: "较上月下降2%"
  }
]

const recentActivities = [
  { id: 1, action: "导入期中考试成绩", time: "2小时前", type: "import" },
  { id: 2, action: "生成学生评语", time: "5小时前", type: "generate" },
  { id: 3, action: "添加学生表现记录", time: "1天前", type: "record" },
  { id: 4, action: "导出成绩报告", time: "2天前", type: "export" }
]

export default function Dashboard() {
  return (
    <div className="space-y-8">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">仪表盘</h1>
          <p className="text-muted-foreground mt-1">
            欢迎使用教师助手，这里是您的教学数据概览
          </p>
        </div>
        <div className="flex items-center gap-2 text-sm text-muted-foreground">
          <Calendar className="h-4 w-4" />
          <span>{new Date().toLocaleDateString('zh-CN', {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            weekday: 'long'
          })}</span>
        </div>
      </div>

      {/* Stats Grid */}
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        {stats.map((stat, index) => (
          <Card key={index} className="hover:shadow-md transition-shadow">
            <CardHeader className="flex flex-row items-center justify-between pb-2">
              <CardTitle className="text-sm font-medium text-muted-foreground">
                {stat.title}
              </CardTitle>
              <stat.icon className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">{stat.value}</div>
              <div className="flex items-center gap-1 mt-1">
                {stat.trend === "up" ? (
                  <ArrowUpRight className="h-3 w-3 text-green-600" />
                ) : (
                  <ArrowDownRight className="h-3 w-3 text-red-600" />
                )}
                <span className={`text-xs ${stat.trend === "up" ? "text-green-600" : "text-red-600"}`}>
                  {stat.change}
                </span>
                <span className="text-xs text-muted-foreground">{stat.description}</span>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>

      {/* Main Content Grid */}
      <div className="grid gap-4 md:grid-cols-7">
        {/* Left Column - Quick Actions */}
        <div className="md:col-span-3 space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>快捷操作</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3">
              <Button className="w-full justify-start" size="lg">
                <GraduationCap className="mr-2 h-4 w-4" />
                导入成绩
              </Button>
              <Button className="w-full justify-start" variant="outline" size="lg">
                <BookOpen className="mr-2 h-4 w-4" />
                记录表现
              </Button>
              <Button className="w-full justify-start" variant="outline" size="lg">
                <Award className="mr-2 h-4 w-4" />
                生成评语
              </Button>
              <Button className="w-full justify-start" variant="outline" size="lg">
                <Users className="mr-2 h-4 w-4" />
                管理学生
              </Button>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>学科成绩分布</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                <div className="flex items-center justify-between text-sm">
                  <span>语文</span>
                  <span className="font-medium">85.2</span>
                </div>
                <Progress value={85.2} className="h-2" />
              </div>
              <div className="space-y-2">
                <div className="flex items-center justify-between text-sm">
                  <span>数学</span>
                  <span className="font-medium">78.5</span>
                </div>
                <Progress value={78.5} className="h-2" />
              </div>
              <div className="space-y-2">
                <div className="flex items-center justify-between text-sm">
                  <span>英语</span>
                  <span className="font-medium">82.1</span>
                </div>
                <Progress value={82.1} className="h-2" />
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Right Column - Recent Activity */}
        <div className="md:col-span-4">
          <Card className="h-full">
            <CardHeader>
              <CardTitle>最近活动</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {recentActivities.map((activity) => (
                  <div key={activity.id} className="flex items-start gap-4 pb-4 border-b last:border-0 last:pb-0">
                    <div className="flex h-8 w-8 items-center justify-center rounded-full bg-muted">
                      <Clock className="h-4 w-4" />
                    </div>
                    <div className="flex-1 space-y-1">
                      <p className="text-sm font-medium">{activity.action}</p>
                      <p className="text-xs text-muted-foreground">{activity.time}</p>
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
