div p-0 id=top
  {{ template "navbar" . }}
  {{ $item := .item }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div mt-6 text-2xl text-center
        div bg-purple-500 rounded-lg p-3
          HomeDucky.com - One Product a Day
        div bg-indigo-100 rounded-lg p-3 text-center
          p
            {{ $item.scheduled_for_human }}
          p id=vote-total-{{$item.asin}}
            {{ $item.votes }}
          a href=/ btn btn-accent btn-sm voter id=vote-{{$item.asin}}
            Vote Up
      div mt-6 flex justify-center text-2xl text-center
        <a href="https://www.amazon.com/dp/{{$item.asin}}?tag=homeduckydotc-20"><img src="{{$item.photo}}"/></a>
      div mt-3 mb-64 space-y-3
        p
          {{ $item.original_title }}
        p text-center
          <a class="btn btn-primary" href="https://www.amazon.com/dp/{{$item.asin}}?tag=homeduckydotc-20">Buy Now</a>
        p text-center
          <a class="btn btn-secondary" href="/core/about">About Us</a>
